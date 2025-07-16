GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=./sql/migrations
GOOSE_DBSTRING=postgres://root:mysecretpassword@localhost:5433/user_service


APP_NAME := server

dev:
	go run ./cmd/$(APP_NAME)

run:
	docker compose up -d && go run go run ./$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

genSqlc:
	sqlc generate

.PHONY: run upse downse resetse

.PHONY: air