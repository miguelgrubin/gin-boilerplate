#!/bin/bash
# Coverage report generator
# Generates test coverage for the pkg directory

go test ./pkg/... -coverprofile=coverage.out.tmp ./pkg/...
cat coverage.out.tmp | grep --invert-match "/mocks/" > coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
