ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=postgres password=test dbname=test host=localhost port=5432 sslmode=disable
endif

ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test_db host=localhost port=5438 sslmode=disable
endif
MIGRATION_FOLDER=C:\Users\Artyom\gitlabozon\homework-5\internal\pkg\db\migrations

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down
	docker-compose down

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)"  create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: test-unit
test-unit:
	cd .\internal\pkg\server
	go test -tags=unit


.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down