APP_NAME ?= slugifyer
VERSION ?= latest

# Image URL to use all building/pushing image targets
IMG ?= "slugifyer:${VERSION}"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: ## Run golangci-lint run
	golangci-lint run

.PHONY: test
test: ## Run unit tests
	go test ./...

.PHONY: build
build: fmt vet ## Build slugifyer binary.
	go build -o bin/slugifyer .


.PHONY: docker-build
docker-build:
	docker build --build-arg="GITHUB_TOKEN=${GITHUB_TOKEN}" -t ${IMG} .

.PHONY: run
run: docker-build
	docker-compose up
