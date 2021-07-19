.PHONY: help all test build build_macos cover fmt

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

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.VcsRef=${VCS_REF} -X main.BuildTime=${BUILD_DATE}"

# ==================================================================
# Build targets
# ==================================================================
default: help

help:
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)
	@echo

all: clean vet test build ## Run tests and build a binary

#vars:
# $(info INFO: Running: vars)
# @echo "HOSTNAME=$(HOSTNAME)"
# @echo "GOROOT=$(GOROOT)"
# @echo "GOPATH=$(GOPATH)"
# @echo

deps: ## Install Go dependencies
	go get 


clean: ## Cleanup and remove artifacts
	GOFLAGS=-mod=mod ${GO} clean -x
	rm -rf \
		./bin \
		${COVPATH}
	@echo

build: vet ## Build binary for Linux
	$(info INFO: Running: build)
	mkdir -p bin/
	GOOS=linux ${GO} build ${LDFLAGS} -o bin/$(PACKAGE) ./cmd/
	@echo

build_macos: ## Build binary for Darwin (macOS)
	mkdir -p bin/
	GOOS=darwin ${GO} build ${LDFLAGS} -o bin/$(PACKAGE) ./cmd/
	@echo

build_win: vars ## Build binary for Windows
	mkdir -p bin/
	GOOS=windows ${GO} build ${LDFLAGS} -o bin/$(PACKAGE) ./cmd/
	@echo

coverage: vet
	${GO} test -count=1 -coverprofile=coverage.out -covermode=count ./...
	@echo

cover: vet check ## Run tests with coverage report html format
	${GO} test -count=1 -coverprofile=coverage.out -covermode=count ./...
	${GO} tool cover -html coverage.out
	@echo

test: vet ## Run tests
	$(info INFO: Running: test)
	${GO} test -race -v ./...
	@echo

vet: ## Run Go vet
	$(info INFO: Running: vet)
	${GO} vet ./...
	@echo

check: ## Run staticcheck code analyzer
	$(info INFO: Running check)
	staticcheck ./...
	@echo
