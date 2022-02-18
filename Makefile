PROJECT_NAME := butler
MODULE_NAME := github.com/butlerhq/$(PROJECT_NAME)
BIN := $(CURDIR)/bin
BIN_AIR := $(CURDIR)/tmp
CMD := $(CURDIR)/cmd
PKG_LIST := $(shell go list ./... | grep -v vendor)
CMD_PKG := $(shell go list $(CMD)/... | grep -v vendor)
COV_PKG := $(shell go list ./... | grep -v gen | tr '\n' ',')
TOOLS := $(CURDIR)/tools
DASHBOARD_DIR := $(CURDIR)/butler-dashboard

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

DOCKER_PUSH ?= false
DOCKER_REGISTRY ?= butlerhq
DOCKER_IMAGE_TAG ?= test
DOCKER_BUILD_FLAGS ?= --progress plain

# Git
GIT_CURRENT_SHA=$(shell git rev-parse --short HEAD)

# Docker compose
DOCKER_COMPOSE_ENV = COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_COMPOSE_CMD = docker-compose -p $(PROJECT_NAME)
DOCKER_COMPOSE_CMD_TEST = $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.test.yml
DOCKER_COMPOSE_CMD_TEST_LOCAL = $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.local.test.yml

# Open Api
OPEN_API_OUT=webapp/config/openapi
OPEN_API_NAME=api
OPEN_API_FILE=$(OPEN_API_OUT)/$(OPEN_API_NAME).swagger.json

# Protobuf
PROTO_ROOT := proto
PROTO_FILES := $(shell find ./proto -name "*.proto")
PROTO_DIRS := $(sort $(dir $(PROTO_FILES)))
PROTO_IMPORTS := -I=$(PROTO_ROOT) -I./vendor/github.com/grpc-ecosystem/grpc-gateway/v2 -I./vendor/github.com/envoyproxy
PROTO_OUT := api
PROTO_CMD := protoc $(PROTO_IMPORTS)

##### Proto #####
proto-all: proto open-api

.PHONY: proto
proto:
	@mkdir -p $(PROTO_OUT)
    # Run protoc separately for each directory because of different package names.
	@for PROTO_FILE in $(PROTO_FILES); do \
		protoc $(PROTO_IMPORTS) \
		 	--go_out=module=$(MODULE_NAME):. --go-grpc_out=module=$(MODULE_NAME):. \
            --grpc-gateway_out . --grpc-gateway_opt module=$(MODULE_NAME) --grpc-gateway_opt logtostderr=true \
			$${PROTO_FILE} && echo "✅ $${PROTO_FILE}" || (echo "❌ $${PROTO_FILE}"; exit 1); \
	done

.PHONY: open-api
open-api:
	@mkdir -p $(OPEN_API_OUT)
	@echo $(PROTO_FILES)
	protoc $(PROTO_IMPORTS) \
    	--openapiv2_out=openapi_naming_strategy=fqn,allow_merge=true,merge_file_name=$(OPEN_API_NAME),logtostderr=true:$(OPEN_API_OUT) \
    	$(PROTO_FILES)

##### Migrations #####
fix-migrations:
	@echo "Fixing migrations"
	@for NAME in $(sort $(dir $(wildcard $(CURDIR)/services/*/migrations/.))); do \
  		echo "Fixing $${NAME}"; \
  		cd $${NAME} && goose fix && echo "All migrations looks good ✅" || (echo "Unable to fix migrations ❌"; exit 1);  \
  	done

##### Binaries #####
tools: clean-tools-bins butler-victorinox
services: clean-services-bins butler-users butler-octopus butler-gateway

clean-tools-bins:
	@echo "Delete old binaries..."
	@rm -f $(BIN)/butler-victorinox

clean-services-bins:
	@echo "Delete old binaries..."
	@rm -f $(BIN)/butler-users
	@rm -f $(BIN)/butler-gateway

butler-victorinox:
	@printf "Build butler-victorinox with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN)/butler-victorinox cmd/victorinox/main.go

butler-gateway:
	@printf "Build butler-gateway service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN)/butler-gateway cmd/gateway/main.go

butler-users:
	@printf "Build butler-users service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN)/butler-users cmd/users/main.go

butler-octopus:
	@printf "Build butler-octopus service with OS: $(GOOS), ARCH: $(GOARCH)..."
	@mkdir -p $(BIN)
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(BIN)/butler-octopus cmd/octopus/main.go


##### Docker #####
docker-all: docker-service-gateway docker-service-users docker-service-octopus docker-webapp docker-victorinox
docker-services: docker-service-gateway docker-service-users docker-service-octopus
docker-tools: docker-victorinox

.PHONY: docker-victorinox
docker-victorinox:
	@printf "Building docker image  $(DOCKER_REGISTRY)/butler-victorinox:$(DOCKER_IMAGE_TAG)...\n"
	docker build . -t $(DOCKER_REGISTRY)/butler-victorinox:$(DOCKER_IMAGE_TAG) $(DOCKER_BUILD_FLAGS) --target victorinox
	@if [ $(DOCKER_PUSH) = true ]; then \
  		echo "Pushing docker image  $(DOCKER_REGISTRY)/butler-victorinox:$(DOCKER_IMAGE_TAG)...\n"; \
		docker push $(DOCKER_REGISTRY)/butler-victorinox:$(DOCKER_IMAGE_TAG); \
	fi

.PHONY: docker-webapp
docker-webapp:
	@printf "Building docker image $(DOCKER_REGISTRY)/butler-webapp:$(DOCKER_IMAGE_TAG)...\n"
	cd ./webapp && docker build . -t $(DOCKER_REGISTRY)/butler-webapp:$(DOCKER_IMAGE_TAG) $(DOCKER_BUILD_FLAGS) --target prod
	@if [ $(DOCKER_PUSH) = true ]; then \
		echo "Pushing docker image  $(DOCKER_REGISTRY)/butler-webapp:$(DOCKER_IMAGE_TAG)...\n"; \
		docker push $(DOCKER_REGISTRY)/butler-webapp:$(DOCKER_IMAGE_TAG); \
	fi
	@cd ..


.PHONY: docker-service-gateway
docker-service-gateway:
	@printf "Building docker image $(DOCKER_REGISTRY)/butler-gateway:$(DOCKER_IMAGE_TAG)...\n"
	@docker build . -t $(DOCKER_REGISTRY)/butler-gateway:$(DOCKER_IMAGE_TAG) $(DOCKER_BUILD_FLAGS) --target service-gateway
	@if [ $(DOCKER_PUSH) = true ]; then \
		echo "Pushing docker image  $(DOCKER_REGISTRY)/butler-gateway:$(DOCKER_IMAGE_TAG)...\n"; \
		docker push $(DOCKER_REGISTRY)/butler-gateway:$(DOCKER_IMAGE_TAG); \
	fi

.PHONY: docker-service-users
docker-service-users:
	@printf "Building docker image $(DOCKER_REGISTRY)/butler-users:$(DOCKER_IMAGE_TAG)...\n"
	@docker build . -t $(DOCKER_REGISTRY)/butler-users:$(DOCKER_IMAGE_TAG) $(DOCKER_BUILD_FLAGS) --target service-users
	@if [ $(DOCKER_PUSH) = true ]; then \
		echo "Pushing docker image  $(DOCKER_REGISTRY)/butler-users:$(DOCKER_IMAGE_TAG)...\n"; \
		docker push $(DOCKER_REGISTRY)/butler-users:$(DOCKER_IMAGE_TAG); \
	fi

.PHONY: docker-service-octopus
docker-service-octopus:
	@printf "Building docker image $(DOCKER_REGISTRY)/butler-octopus:$(DOCKER_IMAGE_TAG)...\n"
	@docker build . -t $(DOCKER_REGISTRY)/butler-octopus:$(DOCKER_IMAGE_TAG) $(DOCKER_BUILD_FLAGS) --target service-octopus
	@if [ $(DOCKER_PUSH) = true ]; then \
		echo "Pushing docker image  $(DOCKER_REGISTRY)/butler-octopus:$(DOCKER_IMAGE_TAG)...\n"; \
		docker push $(DOCKER_REGISTRY)/butler-octopus:$(DOCKER_IMAGE_TAG); \
	fi



lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

.PHONY: vendor
vendor: ## Vendor dependencies
	@echo "Running go.mod vendor"
	@go mod vendor

tidy: ## Clean go.mod dependencies
	@echo "Running go.mod tidy..."
	@go mod tidy

#tools: vendor ## Download and install dependencies
#	@echo "Download and install dependencies, tools..."
#	 cat $(TOOLS)/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install % \
#    		&& echo "✅ Tools installed" || (echo "❌ Failed to install tools"; exit 1);



########################
###  Docker Compose  ###
########################

dev.infra: ## Start dev environment with docker
	@echo "Starting dev infra..."
	@$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.infra.yml up --build postgres redis minio

dev.airbyte: ## Start airbyte
	@echo "Starting airbyte..."
	@$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/airbyte/docker-compose.yml --env-file $(DOCKER_COMPOSE)/airbyte/.env.dev up

dev.migrate: ## Provision databases
	@echo "Starting victorinox..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.services.yml up --build victorinox

.PHONY: dev.services
dev.services: ## Start services with docker in dev environment
	@echo "Starting dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.services.yml up --build users octopus gateway


dev.monitor: ## Start monitor dev evironment with docker
	@echo "Starting monitoring dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.monitor.yml up --build pgadmin

dev.clean: ## Clean docker dev evironment
	@echo "Cleaning dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/airbyte/docker-compose.yml down $(DOCKER_COMPOSE_CLEAN_FLAGS)
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/airbyte/docker-compose.yml rm -f
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.services.yml down $(DOCKER_COMPOSE_CLEAN_FLAGS)
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.services.yml rm -f
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.monitor.yml down $(DOCKER_COMPOSE_CLEAN_FLAGS)
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.monitor.yml rm -f
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.infra.yml down $(DOCKER_COMPOSE_CLEAN_FLAGS)
	$(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.infra.yml rm -f


########################
###       K8s        ###
########################


help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+(\.[a-zA-Z_-]+)*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'



