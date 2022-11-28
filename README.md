# Unit Testing Examples with Mock Interface

## Gomock
For more information & installation instructions on package demonstrated in this repository:
https://github.com/golang/mock

The following make commands can be run in sequence to get started.


## Install dev dependencies
    make install-deps

## Generate mocks for user.go
    mockgen -source=user/service.go -destination=user/mocks/service.go

## Generate all mocks in this repo
    make generate

## Start dev dependencies (postgres) with docker-compose
    make dev

## Migrate database
    make migrate-up

## Start app with live reload
    make start

## Run linter
    make lint

## Run unit tests 
    make test