# Add command to build by operative system
SHELL=/usr/bin/env bash
PROJECTNAME=$(shell basename "$(PWD)")
PWD_PROJECT=$(shell pwd)
LDFLAGS="-X 'main.buildTime=$(shell date)' -X 'main.lastCommit=$(shell git rev-parse HEAD)' -X 'main.semanticVersion=$(shell git describe --tags --dirty=-dev)'"
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)


## help: Get more info on make commands.
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

## build-linux-amd64: Build gofetch for linux amd64
## build gofetch
build:
	@echo "--> Building gofetch binary for $(GOOS):$(GOARCH)"
	@env GOOS=$(GOOS) GOACH=$(GOARCH) go build -o gofetch ./cmd/main.go
	@echo "--> gofetch for $(GOOS):$(GOARCH) built at $(PWD_PROJECT)"

.PHONY: build

