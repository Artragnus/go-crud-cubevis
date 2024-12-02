include .env
export $(shell sed 's/=.*//' .env)

createmigration:
	migrate create -ext sql=sql -dir=sql/migrations -seq init

migrate:
	migrate -path sql/migrations -database ${DATA_SOURCE_NAME} -verbose up

migratedown:
	migrate -path sql/migrations -database ${DATA_SOURCE_NAME} -verbose down

.PHONY: migrate createmigration migratedown