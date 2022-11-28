# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# --- Tooling & Variables ----------------------------------------------------------------

include ./misc/make/vars.Makefile
include ./misc/make/tools.Makefile
include ./misc/make/help.Makefile

# --- Development Environment ------------------------------------------------------------

install-deps: migrate air gotestsum tparse mockgen ## Install Development Dependencies (localy).
deps: $(MIGRATE) $(AIR) $(GOTESTSUM) $(TPARSE) $(MOCKGEN) ## Checks for Global Development Dependencies.
deps:
	@echo "Required Tools Are Available"

dev:
	docker-compose up -d
	docker container ls
	docker-compose ls

down:
	docker-compose down

start: $(AIR)
	air -c .air.toml

# --- Code Actions ------------------------------------------------------------------------

generate: $(MOCKGEN)
	@mockgen -version
	go generate -v ./...

lint: $(GOLANGCI) ## Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

TESTS_ARGS := --format testname --jsonfile gotestsum.json.out
TESTS_ARGS += --max-fails 2
TESTS_ARGS += -- ./...
TESTS_ARGS += -test.parallel 2
TESTS_ARGS += -test.count    1
TESTS_ARGS += -test.failfast
TESTS_ARGS += -test.coverprofile   coverage.out
TESTS_ARGS += -test.timeout        5s
TESTS_ARGS += -race

test: $(GOTESTSUM) $(TPARSE) ## Run Tests & parse details
	@gotestsum $(TESTS_ARGS)
	@cat gotestsum.json.out | $(TPARSE) -all -top -notests

# --- Database Migrations ----------------------------------------------------------------

PG_DSN := "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"

migrate-up: $(MIGRATE) ## Apply all (or N up) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(PG_DSN) -path=postgres/migrations up ${NN}

.PHONY: migrate-down
migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(PG_DSN) -path=postgres/migrations down ${NN}

.PHONY: migrate-drop
migrate-drop: $(MIGRATE) ## Drop everything inside the database.
	migrate  -database $(PG_DSN) -path=postgres/migrations drop

.PHONY: migrate-create
migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@ read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir postgres/migrations $${Name}