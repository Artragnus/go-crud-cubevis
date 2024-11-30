FROM golang:1.22 as BUILD

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM scratch

WORKDIR /app

COPY --from=BUILD /app/api ./

COPY --from=BUILD /app/.env ./

EXPOSE 8080

CMD ["./api"]