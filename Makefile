#check if exist .env file
ifneq (,$(wildcard .env))
include .env
export
endif

.PHONY: compose-up compose-logs compose-down build swag tidy db-psql db-reset

compose-up:
	docker compose up -d --build

compose-logs:
	docker compose logs -f

compose-down:
	docker compose down -v

test-api:
	bruno run bruno_collection/

db-psql:
	docker exec -it events-db psql -U $(DB_USER) -d $(DB_NAME)

db-reset: compose-down compose-up

tidy:
	go mod tidy

swag:
	swag init -g cmd/app/main.go

build:
	docker build -t events-api:local .
