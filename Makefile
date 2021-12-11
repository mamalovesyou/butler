.PHONY: proto

PROJECT_NAME := butler
MODULE_NAME := github.com/butlerhq/$(PROJECT_NAME)
BIN := $(CURDIR)/bin
BIN_AIR := $(CURDIR)/tmp
CMD := $(CURDIR)/cmd
PKG_LIST := $(shell go list ./... | grep -v vendor)
CMD_PKG := $(shell go list $(CMD)/... | grep -v vendor)
COV_PKG := $(shell go list ./... | grep -v gen | tr '\n' ',')
TOOLS := $(CURDIR)/tools
WEBAPP := $(CURDIR)/webapp

AIR_TARGET=tmp/main
TARGET_PKG=$(MODULE_NAME)/${PACKAGE_PATH}


DEPLOYMENT := $(CURDIR)/deployment
DOCKER_COMPOSE := $(DEPLOYMENT)/docker-compose
DOCKER_COMPOSE_CLEAN_FLAGS=--volumes --rmi local --remove-orphans

##### Arguments ######

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(if $(shell go env GOBIN),$(shell go env GOBIN),$(GOPATH)/bin)

# Name resolution requires cgo to be enabled on macOS and Windows: https://golang.org/pkg/net/#hdr-Name_Resolution.
ifndef CGO_ENABLED
	ifeq ($(GOOS),linux)
	CGO_ENABLED := 0
	else
	CGO_ENABLED := 1
	endif
endif

DOCKER_IMAGE_TAG ?= test

# Git
GIT_CURRENT_SHA=$(shell git rev-parse --short HEAD)

# Docker compose
DOCKER_COMPOSE_ENV = COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_COMPOSE_CMD = $(DOCKER_COMPOSE_ENV) docker-compose -p $(PROJECT_NAME)
DOCKER_COMPOSE_CMD_TEST = $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.test.yml
DOCKER_COMPOSE_CMD_TEST_LOCAL = $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.local.test.yml

# Open Api
OPEN_API_OUT=config/openapi
OPEN_API_NAME=api
OPEN_API_FILE=$(OPEN_API_OUT)/$(OPEN_API_NAME).swagger.json

# Protobuf
PROTO_ROOT := proto
PROTO_FILES := $(shell find ./proto/services -name "*.proto")
PROTO_DIRS := $(sort $(dir $(PROTO_FILES)))
PROTO_IMPORTS := -I=$(PROTO_ROOT) -I./vendor/github.com/grpc-ecosystem/grpc-gateway/v2 -I./vendor/github.com/envoyproxy -I./third_party
PROTO_OUT := api
PROTO_CMD := protoc $(PROTO_IMPORTS)

##### Proto #####
proto:
	@mkdir -p $(PROTO_OUT)
    # Run protoc separately for each directory because of different package names.
	@for PROTO_FILE in $(PROTO_FILES); do \
		protoc $(PROTO_IMPORTS) \
		 	--go_out=module=$(MODULE_NAME):. --go-grpc_out=module=$(MODULE_NAME):. \
            --grpc-gateway_out . --grpc-gateway_opt module=$(MODULE_NAME) --grpc-gateway_opt logtostderr=true \
			$${PROTO_FILE} && echo "✅ $${PROTO_FILE}" || (echo "❌ $${PROTO_FILE}"; exit 1); \
	done

open-api:
	@mkdir -p $(OPEN_API_OUT)
	@echo $(PROTO_FILES)
	protoc $(PROTO_IMPORTS) \
    	--openapiv2_out=openapi_naming_strategy=fqn,allow_merge=true,merge_file_name=$(OPEN_API_NAME),logtostderr=true:$(OPEN_API_OUT) \
    	$(PROTO_FILES)

##### Binaries #####

services: clean-bins butler-users

clean-bins:
	@echo "Delete old binaries..."
	@rm -f $(BIN)/butler-users
	@rm -f $(BIN)/butler-gateway

butler-gateway:
	@printf "Build butler-gateway service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN)/butler-gateway cmd/services/gateway/main.go

butler-users:
	@printf "Build butler-users service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN)/butler-users cmd/services/users/main.go


##### Docker #####
docker-service-gateway:
	@printf "Building docker image butlerhq/butler-users:$(DOCKER_IMAGE_TAG)..."
	docker build . -t butlerhq/butler-gateway:$(DOCKER_IMAGE_TAG) --build-arg TARGET=service-gateway

docker-service-users:
	@printf "Building docker image butlerhq/butler-users:$(DOCKER_IMAGE_TAG)..."
	docker build . -t butlerhq/butler-users:$(DOCKER_IMAGE_TAG) --build-arg TARGET=service-users



lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}


tidy: ## Clean go.mod dependencies
	@echo "Running go.mod tidy..."
	@go mod tidy

dependencies: tidy ## Download and install dependencies
	@echo "Download and install dependencies, tools..."
	cat $(TOOLS)/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install % \
    		&& echo "✅ Tools installed" || (echo "❌ Failed to install tools"; exit 1);




# DOCKER ENV
docker.dev.infra: ## Start dev environment with docker
	@echo "Starting dev infra..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.dev.yml up --build --remove-orphans postgres redis jaeger swagger-ui

docker.dev.provision: ## Provision databases
	@echo "Starting victorinox..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.dev.yml up --build --remove-orphans victorinox


docker.dev.services: ## Start services with docker in dev environment
	@echo "Starting dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.dev.yml up --build --abort-on-container-exit --remove-orphans auth workspace gateway


docker.dev.monitor: ## Start monitor dev evironment with docker
	@echo "Starting monitoring dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.dev.yml up --build --remove-orphans pgadmin

docker.dev.clean: ## Clean docker dev evironment
	@echo "Cleaning dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.dev.yml down $(DOCKER_COMPOSE_CLEAN_FLAGS)
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.dev.yml rm -f
	@echo "Cleaning monitor dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.monitor.dev.yml down $(DOCKER_COMPOSE_CLEAN_FLAGS)
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.monitor.dev.yml rm -f


clean: docker.dev.clean ## Clean all
	@echo "Cleaning ..."
	rm -rf $(BIN)
	rm -rf $(CURDIR)/tmp

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+(\.[a-zA-Z_-]+)*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'