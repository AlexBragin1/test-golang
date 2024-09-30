# check if .env exists
ifneq (,$(wildcard .env))
	include .env
	export
endif

PG_URL="postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

MG_PATH="test-golang/db/migrations/"

.PHONY: generate
generate:
	go generate ./...



.PHONY: run
run:
	go run main.go

.PHONY: up
up:
	docker compose up --build -d app db

.PHONY: infra
infra:
	docker compose up -d --build db


.PHONY: vet
vet:
	go vet ./...

.PHONY: down
down:
	docker compose down

.PHONY: ps
ps:
	docker compose ps

.PHONY: migrate_up
test_migrate_up:
	migrate -path $(MG_PATH) -database $(PG_TEST_URL) -verbose up

.PHONY: migrate_up
migrate_up:
	migrate -path $(MG_PATH) -database $(PG_URL) -verbose up

.PHONY: migrate_down
migrate_down:
	migrate -path $(MG_PATH) -database $(PG_URL) -verbose down

.PHONY: docker_migrate_up
docker_migrate_up:
	docker compose exec app \
		migrate -path $(MG_PATH) -database $(PG_URL) -verbose up

.PHONY: docker_migrate_down
docker_migrate_down:
	docker compose exec app \
		migrate -path $(MG_PATH) -database $(PG_URL) -verbose down


.PHONY: create_migration
create_migration:
	migrate create -ext sql -dir $(MG_PATH) -seq users


