# Test Go

## Running testings

- run all test `go test ./...`
  
- run all test with coverage `go test -cover ./...`

- run all test with coverage and create file coverage `go test -coverprofile=coverage.out ./...`

- open file html with file coverage `go tool cover -html=coverage.out`

## Automatization commands with Makefile

```
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
```