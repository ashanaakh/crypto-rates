BIN ?=	crates
PKG :=	$(shell basename $(CURDIR))

default: run

.PHONY: run
run: build
	@./$(BIN)

.PHONY: build
build:
	@go build -o $(BIN)

.PHONY: install
install:
	@echo "Installing $(PKG)..."
	@go install

.PHONY: check
check: lint vet

.PHONY: list
lint:
	@echo "Linting your go code..."
	@golint -set_exit_status $(go list ./...)

.PHONY: vet
vet:
	@echo "Vet in progress..."
	@go vet ./...

.PHONY: list
list:
	@go list ./...

.PHONY: deps
deps:
	@go get ./...
