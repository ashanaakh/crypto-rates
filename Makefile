BIN ?= rates

default: run

.PHONY: build
build:
	@go build -o $(BIN)

.PHONY: build
run: build
	@./$(BIN)

.PHONY: install
install:
	@go install
