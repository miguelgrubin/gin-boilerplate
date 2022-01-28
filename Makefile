.PHONY: help
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
	cd bin && ./app create config
	cd bin && ./app serve
lint:                         ## Show lint errors
	golangci-lint run
test:                         ## Run all tests
	go test ./...
test-server:                  ## Start testing server
	goconvey
