FROM golang:1.22 as BUILD

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM scratch
FROM alpine:latest

RUN apk --no-cache add curl bash make && \
    curl -sSL https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar -xz && \
    mv migrate /usr/local/bin/

WORKDIR /app

COPY --from=BUILD /app/api ./
COPY --from=BUILD /app/.env ./

COPY --from=BUILD /app/sql/migrations/ ./sql/migrations
COPY --from=BUILD /app/Makefile ./

EXPOSE ${PORT}

CMD ["bash", "-c", "make migrate && ./api"]