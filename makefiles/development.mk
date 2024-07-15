PROJECT_PATH := $(shell pwd)/src/umom

PHONY: tidy ## Make go mod tidy to everything recursively
tidy:
	@cd $(PROJECT_PATH) && go get -u && go mod tidy

PHONY: test ## Run tests
test:
	@cd $(PROJECT_PATH) && go test -v ./...

PHONY: run ## Run the project. Example: make run ARGS=$HOME/Desktop/test/
run:
	@cd $(PROJECT_PATH) && go run main.go $(ARGS)
