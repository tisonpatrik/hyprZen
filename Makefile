tidy:
	go mod tidy -v
	go fmt ./...

build:
	@go get -u ./...
	@go build -o build/main ./cmd

run: build
	./build/main

.PHONY: run tidy build
