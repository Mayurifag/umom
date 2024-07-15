PROJECT_PATH := $(shell pwd)/src/umom

PHONY: tidy ## Make go mod tidy to everything recursively
tidy:
	@cd $(PROJECT_PATH) && go get -u && go mod tidy

PHONY: test ## Run tests
test:
	@cd $(PROJECT_PATH) && go test -v ./...
