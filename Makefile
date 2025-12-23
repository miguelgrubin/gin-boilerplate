.PHONY: help all test test/cover clean

help:         ## Show this help.
	@echo "Makefile commnads list\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
install:      ## Install dependencies with go mod
	go mod vendor
build:        ## Clean & Compile
	rm -f bin/*
	go build -o bin/app -race
run: build
	cd bin && ./app create-config
	cd bin && ./app serve
lint:         ## Show lint errors
	revive -config revive.toml -exclude vendor/... -formatter friendly ./...
sec:          ## Show security errors
	gosec ./...
format:       ## Apply format to all files
	gofmt -s -w .
test:         ## Run all tests
	go test ./pkg/... -v
coverage:     ## Run all tests and show coverage
	./scripts/coverage.sh
start:        ## Run hot-reload server
	air
debug:        ## Run server on debug mode
	dlv debug
