# Add command to build by operative system
SHELL=/usr/bin/env bash
PROJECTNAME=$(shell basename "$(PWD)")
PWD_PROJECT=$(shell pwd)
LDFLAGS="-X 'main.buildTime=$(shell date)' -X 'main.lastCommit=$(shell git rev-parse HEAD)' -X 'main.semanticVersion=$(shell git describe --tags --dirty=-dev)'"
GOOS_LINUX="linux"
GOOS_MAC="darwin"
GOOS_WINDOWS="windows"

GOARCH_AMD64="amd64"
GOARCH_ARM="arm"
GOARCH_ARM64="arm64"

## help: Get more info on make commands.
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

## build-linux-amd64: Build gofetch for linux amd64
build-linux-amd64:
	@echo "--> Building gofetch binary for $(GOOS_LINUX):$(GOARCH_AMD64)"
	@env GOOS=$(GOOS_LINUX) GOACH=$(GOARCH_AMD64) go build -o gofetch ./cmd/linux
	@echo "--> gofetch for $(GOOS_LINUX):$(GOARCH_AMD64) built at $(PWD_PROJECT)"
.PHONY: build-linux-amd64

## build-linux-arm: Build gofetch for linux arm
build-linux-arm:
	@echo "--> Building gofetch binary for $(GOOS_LINUX):$(GOARCH_ARM)"
	@env GOOS=$(GOOS_LINUX) GOACH=$(GOARCH_ARM) go build -o gofetch ./cmd/linux
	@echo "--> gofetch for $(GOOS_LINUX):$(GOARCH_ARM) built at $(PWD_PROJECT)"
.PHONY: build-linux-arm

## build-linux-arm64: Build gofetch for linux arm64
build-linux-arm64:
	@echo "--> Building gofetch binary for $(GOOS_LINUX):$(GOARCH_ARM64)"
	@env GOOS=$(GOOS_LINUX) GOACH=$(GOARCH_ARM64) go build -o gofetch ./cmd/linux
	@echo "--> gofetch for $(GOOS_LINUX):$(GOARCH_ARM64) built at $(PWD_PROJECT)"
.PHONY: build-linux-arm64

## build-mac-amd64: Build gofetch for mac amd64
build-mac-amd64:
	@echo "--> Building gofetch binary for $(GOOS_MAC):$(GOARCH_AMD64)"
	@env GOOS=$(GOOS_MAC) GOACH=$(GOOS_MAC) go build -o gofetch ./cmd/mac
	@echo "--> gofetch for $(GOOS_MAC):$(GOARCH_AMD64) built at $(PWD_PROJECT)"
.PHONY: build-mac-amd64

## build-mac-arm: Build gofetch for mac arm
build-mac-arm:
	@echo "--> Building gofetch binary for $(GOOS_MAC):$(GOARCH_ARM)"
	@env GOOS=$(GOOS_MAC) GOACH=$(GOOS_MAC) go build -o gofetch ./cmd/mac
	@echo "--> gofetch for $(GOOS_MAC):$(GOARCH_ARM) built at $(PWD_PROJECT)"
.PHONY: build-mac-arm

## build-mac-arm64: Build gofetch for mac arm64
build-mac-arm64:
	@echo "--> Building gofetch binary for $(GOOS_MAC):$(GOARCH_ARM64)"
	@env GOOS=$(GOOS_MAC) GOACH=$(GOOS_MAC) go build -o gofetch ./cmd/mac
	@echo "--> gofetch for $(GOOS_MAC):$(GOARCH_ARM64) built at $(PWD_PROJECT)"
.PHONY: build-mac-arm64

## build-windows-amd64: Build gofetch for windows amd64
build-windows-amd64:
	@echo "--> Building gofetch binary for $(windows):$(GOARCH_AMD64)"
	@env GOOS=$(windows) GOACH=$(windows) go build -o gofetch ./cmd/mac
	@echo "--> gofetch for $(windows):$(GOARCH_AMD64) built at $(PWD_PROJECT)"
.PHONY: build-windows-amd64


