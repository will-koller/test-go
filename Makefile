#!make
MAKEFLAGS += --silent

test-all:
	go test ./...

test-cover:
	go test -cover ./...

test-coverage:
	go test -coverprofile=coverage.out ./...

read-cover:
	go tool cover -html=coverage.out