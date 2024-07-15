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

PHONY: build ## Build the project
build:
	@cd $(PROJECT_PATH) && go build -o ../../bin/umom main.go

# TODO: only if the folder in $PATH
PHONY: copy-to-usr-bin ## Copy the binary to ~/.local/bin
copy-to-usr-bin:
	chmod +x ./bin/umom
	cp ./bin/umom ~/.local/bin

PHONY: build-and-install ## Build and install the project into ~/.local/bin
build-and-install: tidy test build copy-to-usr-bin
