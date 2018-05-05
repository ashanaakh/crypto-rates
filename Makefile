BIN ?=	crypto-rates
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

.PHONY: lint
lint:
	@echo "Linting your go code..."
	@golint -set_exit_status $(go list ./...)

.PHONY: vet
vet:
	@echo "Vet in progress..."
	@go vet ./...

.PHONY: deps
deps:
	@go get ./...

.PHONY: test
test:
	@go test -v ./...

