PROJECT_NAME := butler
MODULE_NAME := github.com/matthieuberger/$(PROJECT_NAME)
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


# Dockerfiles
API_DOCKERFILE=build/api/Dockerfile

# Git
GIT_CURRENT_SHA=$(shell git rev-parse --short HEAD)

# Docker compose
DOCKER_COMPOSE_ENV = COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_COMPOSE_CMD = $(DOCKER_COMPOSE_ENV) docker-compose -p $(PROJECT_NAME)
DOCKER_COMPOSE_CMD_TEST = $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.test.yml
DOCKER_COMPOSE_CMD_TEST_LOCAL = $(DOCKER_COMPOSE_CMD) -f $(DOCKER_COMPOSE)/docker-compose.local.test.yml

# Protobuf
PROTO_DIR=api/proto
PROTO_GO_OUT=internal/api
PROTOC_DIR_OPTS = -I./vendor/github.com/grpc-ecosystem/grpc-gateway/v2 -I./vendor/github.com/envoyproxy -I./third_party -I.
PROTO_CMD = protoc $(PROTOC_DIR_OPTS)
PROTO_GTW_FILE = $(PROTO_DIR)/gateway.proto
PROTO_TS_OUT_DIR = $(WEBAPP)/api
PROTOC_GEN_TS_PATH = $(WEBAPP)/node_modules/.bin/protoc-gen-ts


# OpenApi
OPEN_API_DIR=api/openapi
OPEN_API_NAME=apiV1
OPEN_API_FILE=$(OPEN_API_DIR)/$(OPEN_API_NAME).swagger.json


all: build.all

# TOOLS

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

race: vendor ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: vendor ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

vendor: ## Vendor go.mod dependencies
	@echo "Download vendor dependencies..."
	go mod vendor

tidy: ## Clean go.mod dependencies
	@echo "Cleaning go.mod dependencies..."
	@go mod tidy

tools: vendor ## Install tools
	@echo "Installing tools..."
	@cat $(TOOLS)/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install % \
		&& echo "✅ Tools installed" || (echo "❌ Failed to install tools"; exit 1);

air: ## Build air for developpment
	time go install github.com/cosmtrek/air

install.victorinox: ## Build api for production
	time go install $(CMD)/victorinox


# BUILD CMD
build.dev: ## Build api for Air (Live reloading)
	time go build -i -o $(AIR_TARGET) ${TARGET_PKG}


# Run tests TEST
unit: ## Run unit tests
	@go test -short -coverpkg=$(COV_PKG) -coverprofile=.coverage.txt ./...

e2e: ## Run e2e tests
	@echo "Running e2e tests..."
	@go test -v ./test/e2e/... --tags=e2e


# Run test environment
test.e2e.infra.start: ## Start e2e tests evironment with docker-compose
	@echo "Starting e2e tests env..."
	@docker-compose -f $(DOCKER_COMPOSE)/docker-compose.test.yml up --build --no-deps postgres-test api-test

test.e2e.start: ## Run e2e tests
	@echo "Running e2e tests..."
	@docker-compose -f $(DOCKER_COMPOSE)/docker-compose.test.yml up --build --no-deps e2e-test

test.e2e.clean: ## Clean e2e tests env
	@docker-compose -f $(DOCKER_COMPOSE)/docker-compose.test.yml down

test.e2e:
	# Build containers
	$(DOCKER_COMPOSE_CMD_TEST) build --progress=plain
	# Run service containers, except for test container
	$(DOCKER_COMPOSE_CMD_TEST) up -d --renew-anon-volumes --remove-orphans --force-recreate postgres-test api-test
	# Run migrations
	$(DOCKER_COMPOSE_CMD_TEST) run --rm migrations
	# Run the tests
	$(DOCKER_COMPOSE_CMD_TEST) run --rm tests
	# Run cleanup
	$(DOCKER_COMPOSE_CMD_TEST) down


test.e2e.local.infra:
	# Build containers
	$(DOCKER_COMPOSE_CMD_TEST_LOCAL) build --progress=plain
	# Run service containers, except for test container
	$(DOCKER_COMPOSE_CMD_TEST_LOCAL) up -d --remove-orphans --renew-anon-volumes --force-recreate postgres-test api-test
	# Run migrations
	$(DOCKER_COMPOSE_CMD_TEST_LOCAL) run --rm migrations

test.e2e.local.start:
	# Run the tests
	$(DOCKER_COMPOSE_CMD_TEST_LOCAL) up --force-recreate tests



# CONTAINERS
container.api.base:
	@echo "Building api container..."
	@DOCKER_BUILDKIT=1 docker build --target base -t $(PROJECT_NAME)-api:latest -t $(PROJECT_NAME)-api:$(GIT_CURRENT_SHA) -f $(API_DOCKERFILE) .

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


# PROTO
# TODO: Add --validate_out="module=$(MODULE_NAME),lang=go:." when module will be fixed
gen.services: $(PROTO_DIR)/* ## Regenerate go files from proto-rest files
	@for file in $^ ; do \
		$(PROTO_CMD) --go_out=module=$(MODULE_NAME):. \
		--go-grpc_out=module=$(MODULE_NAME):. \
		--grpc-gateway_out . \
        --grpc-gateway_opt module=$(MODULE_NAME) \
		--grpc-gateway_opt logtostderr=true \
		$${file} && echo "✅ $${file} generated" || (echo "❌ $${file} failed"; exit 1); \
	done

gen.openapi: $(PROTO_DIR)/* ## Generate openapi doc from proto
	@echo "Generating swagger from proto files..."
	@mkdir -p $(OPEN_API_DIR)
	@$(PROTO_CMD) \
    	--openapiv2_out=fqn_for_openapi_name=true,allow_merge=true,merge_file_name=$(OPEN_API_NAME),logtostderr=true:$(OPEN_API_DIR) $^ \
		&& echo "✅ Succesffuly generated swagger file" || (echo "❌ Failed generated swagger file"; exit 1); \

gen.web.old: $(PROTO_DIR)/* ## Generate openapi specs from proto-rest annotations
	@echo "Generating typescript api client..."
	@mkdir -p $(PROTO_TS_OUT_DIR)
	#$(PROTO_CMD) --plugin="protoc-gen-ts=$(PROTOC_GEN_TS_PATH)" \
#         --js_out="import_style=commonjs,binary:$(PROTO_TS_OUT_DIR)" \
#         --ts_out="service=true:$(PROTO_TS_OUT_DIR)" $^ \
#         && echo "✅ Api client generated" || (echo "❌ Failed to generate api client"; exit 1);
	$(PROTO_CMD) --js_out="import_style=commonjs,binary:$(PROTO_TS_OUT_DIR)" \
		--grpc-web_out=import_style=typescript,mode=grpcweb:$(PROTO_TS_OUT_DIR) protoc-gen-validate/validate/validate.proto $^ \
		&& echo "✅ Api client generated" || (echo "❌ Failed to generate api client"; exit 1);

gen.web: $(PROTO_DIR)/* ## Generate openapi specs from proto-rest annotations
	@echo "Generating typescript api client..."
	@mkdir -p $(PROTO_TS_OUT_DIR)
	$(PROTO_CMD) --plugin=./webapp/node_modules/.bin/protoc-gen-ts_proto \
		--ts_proto_opt=outputServices=false --ts_proto_opt=outputClientImpl=grpc-web \
		--ts_proto_out=$(PROTO_TS_OUT_DIR) $(PROTO_GTW_FILE) \
		&& echo "✅ Api client generated" || (echo "❌ Failed to generate api client"; exit 1);

clean: docker.dev.clean ## Clean all
	@echo "Cleaning ..."
	rm -rf $(BIN)
	rm -rf $(CURDIR)/tmp

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+(\.[a-zA-Z_-]+)*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'