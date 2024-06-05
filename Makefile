COMPOSE_NAME = crime-stats

LOCAL_COMPOSE_FILE ?= docker-compose.yaml
LOCAL_PROJECT_NAME := $(COMPOSE_NAME)-local

.PHONY: build
build:
	docker network inspect service >/dev/null 2>&1 || docker network create --driver bridge service
	docker-compose -p $(LOCAL_PROJECT_NAME) -f $(LOCAL_COMPOSE_FILE) build --progress=plain

.PHONY: up
up:
	POSTGRES_DATA=postgres-data docker-compose -p $(LOCAL_PROJECT_NAME) -f $(LOCAL_COMPOSE_FILE) up

.PHONY: down
down:
	docker-compose -p $(LOCAL_PROJECT_NAME) -f $(LOCAL_COMPOSE_FILE) down

.PHONY: local
local: build up
