SHELL := /bin/bash
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GH_BRANCH := $(shell echo $(BRANCH) | sed 's@.*/@@')

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## run Go tests
	go test ./... -v

deps: ## make sure deps are up to date
	dep ensure

docker-build: ## build Docker image
	docker build -t amm2272/filament:$(GH_BRANCH) .

docker-test: ## build and run tests in Docker image
	docker run amm2272/filament:$(GH_BRANCH) go test ./...

docker-push: ## push Docker image
	docker push amm2272/filament:$(GH_BRANCH)
