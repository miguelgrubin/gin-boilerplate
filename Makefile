.PHONY: help all test test/cover clean

GO=go
GOCOVER=$(GO) tool cover

help:                         ## Show this help.
	@echo "\tMakefile commnads list\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
install:                      ## Install dependencies with go mod
	go mod vendor
compile:                      ## Compile App
	go build -o bin/app -race
clean:                        ## Delete build and configs
	rm -f bin/*
build: clean compile          ## Clean & Compile
run: build
	cd bin && ./app create-config
	cd bin && ./app serve
lint:                         ## Show lint errors
	go vet ./...
format:                       ## Apply format to all files
	gofmt -s -w .
test:                         ## Run all tests
	go test ./... -v
test/cover:                         ## Run all tests
	go test ./... -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
start/dev:
	air
