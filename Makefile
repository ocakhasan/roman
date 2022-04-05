.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## builds the docker image
	@docker build -t ${APP_NAME} .

run: ## creates and runs the container
	@docker-compose up -d adobe

stop: ## stops the container
	@docker-compose stop

down: ## removes the container
	@docker-compose down

test: ## runs test in all of the modules
	@go test -v ./...

