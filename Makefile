.PHONY: help
help:
	echo "Help"
install:
	go mod vendor
compile:
	go build -o bin/app -race
clean:
	rm -f bin/app
run: clean compile
	./bin/app
test:
	go test ./...
test-server:
	goconvey
