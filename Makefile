# Add command to build by operative system

SHELL := /bin/bash
BIN_DIR := $(CURDIR)/bin
PROJECTNAME=$(shell basename "$(PWD)")
PWD_PROJECT=$(shell pwd)
LDFLAGS="-X 'main.buildTime=$(shell date)' -X 'main.lastCommit=$(shell git rev-parse HEAD)' -X 'main.semanticVersion=$(shell git describe --tags --dirty=-dev)'"
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

export GOBIN := $(BIN_DIR)


## help: Get more info on make commands.
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

## build gofetch
build:
	@echo "--> Building gofetch binary for $(GOOS):$(GOARCH)"
	@env CGO_ENABLED=0 GOOS=$(GOOS) GOACH=$(GOARCH) go build -v -installsuffix cgo -o gofetch ./cmd/main.go
	@echo "--> gofetch for $(GOOS):$(GOARCH) built at $(PWD_PROJECT)"

.PHONY: build

## run linter
linter:
	@echo "Checking code"
	$(BIN_DIR)/golangci-lint run $(CURDIR)/...

.PHONY: linter

setup-linter:
	@echo "Installing golanglint dependency"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: setup-linter
