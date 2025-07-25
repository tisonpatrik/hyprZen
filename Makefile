tidy:
	go mod tidy -v
	go fmt ./...

run:
	@go get -u ./...
	@go build -o build/main ./cmd
	./build/main

.PHONY: run tidy
