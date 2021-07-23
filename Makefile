.PHONY: help dep test build build_macos build_win cover

# Grab all make targets comments and print them out.
define PRINT_HELP_PYSCRIPT
import re, sys

for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-20s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

# ==================================================================
# ManHat Setup
# ==================================================================
VERSION := 0.1.0
VCS_REF	:= `git rev-parse HEAD`
BUILD_DATE := `date -u +"%Y-%m-%d-%H-%M-%SZ"`
PACKAGE = manhat

# ==================================================================
# Go Environment Setup
# ==================================================================
export GO111MODULE=on
export GOFLAGS=-mod=vendor

GO=${shell which go}
SHELL :=/bin/bash
ROOT := $(shell pwd)
# List of all .go files in the project, excluding vendor and .tools
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.gopath/*")

# LDFLAGS for building binary for linux, mac, win
IMP_PATH=github.com/qba73/manhat
LDFLAGS=-ldflags "-X ${IMP_PATH}.version=${VERSION} -X ${IMP_PATH}.vcsref=${VCS_REF} -X ${IMP_PATH}.buildtime=${BUILD_DATE}"

# ==================================================================
# Build targets
# ==================================================================
default: help

help:
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)
	@echo

dep: ## Install Go dependencies
	${GO} mod tidy
	${GO} mod verify
	${GO} mod vendor 
	@echo

clean: ## Cleanup and remove artifacts
	GOFLAGS=-mod=mod ${GO} clean -x
	rm -rf \
		./bin \
		${COVPATH}
	@echo

build: vet ## Build binary for Linux
	mkdir -p bin/
	GOOS=linux ${GO} build ${LDFLAGS} -o bin/$(PACKAGE) ./cmd/
	@echo

build_macos: vet ## Build binary for Darwin (macOS)
	mkdir -p bin/
	GOOS=darwin ${GO} build ${LDFLAGS} -o bin/$(PACKAGE) ./cmd/
	@echo

build_win: ## Build binary for Windows
	mkdir -p bin/
	GOOS=windows ${GO} build ${LDFLAGS} -o bin/$(PACKAGE) ./cmd/
	@echo

coverage: vet
	${GO} test -count=1 -coverprofile=coverage.out -covermode=count ./...
	@echo

cover: vet ## Run tests with coverage report html format
	${GO} test -count=1 -coverprofile=coverage.out -covermode=count ./...
	${GO} tool cover -html coverage.out
	@echo

test: ## Run tests
	${GO} test -race 
	@echo

vet: ## Run Go vet
	${GO} vet ./...
	@echo
