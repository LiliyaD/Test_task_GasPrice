.PHONY: run
run: build
	./bin/ethereum

.PHONY: build
build:
	go build -o bin/ethereum cmd/main.go