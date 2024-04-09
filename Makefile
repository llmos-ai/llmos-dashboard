# Directory of Makefile
export ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

REPO?=docker.io/llmosai/llmos-dashboard
GIT_COMMIT?=$(shell git rev-parse HEAD)
GIT_COMMIT_SHORT?=$(shell git rev-parse --short HEAD)
GIT_TAG?=$(shell git describe --candidates=50 --abbrev=0 --tags 2>/dev/null || echo "v0.0.0-dev" )
VERSION?=$(GIT_TAG)

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# CONTAINER_TOOL defines the container tool to be used for building images.
CONTAINER_TOOL ?= docker
BUILDKIT_PROGRESS=plain

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
.PHONY: generate
generate: ## Generate code containing ent schema and manifest in generate.go
	go generate

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.54.2
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter & yamllint
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix

##@ Build
GORELEASE = $(shell pwd)/bin/goreleaser
GORELEASE_VERSION ?= v1.25.1
gorelease:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sfL https://goreleaser.com/static/run | VERSION=${GORELEASE_VERSION} DISTRIBUTION=oss h -s -- check -b $(shell dirname $(GORELEASE)) ;\
	}

.PHONY: build
build: build-ui gorelease ## Build dashboard using goreleaser.
	VERSION=$(VERSION) \
	goreleaser release --snapshot --clean

# brew install FiloSottile/musl-cross/musl-cross is required
.PHONY: darwin-build
darwin-build: build-ui gorelease ## Build dashboard on darwin env using goreleaser.
	VERSION=$(VERSION) \
	goreleaser release -f .goreleaser-darwin.yaml --snapshot --clean

.PHONY: local-build
local-build: ## Build dashboard only using goreleaser
	VERSION=$(VERSION) \
	goreleaser release -f .goreleaser-darwin.yaml --snapshot --clean --verbose

.PHONY: build-ui
build-ui:
	rm -rf $(ROOT_DIR)/ui/build
	$(CONTAINER_TOOL) buildx build --progress=$(BUILDKIT_PROGRESS) \
			-o type=local,dest=$(ROOT_DIR)/ui \
			-t artifact \
			-f package/Dockerfile-ui .

.PHONY: run
run: ## Run the binary
	go run main.go

