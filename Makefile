.PHONY: vendor
vendor: ## run mod tidy and mod vendor
	@go mod tidy && go mod vendor

.PHONY: build
build: ## Build go binary and copy it to a docker container running ubuntu
	@docker build -t web-page-analyzer .

.PHONY: run
run: ## Run application in docker container
	@docker run --rm -it -p 8080:8080 web-page-analyzer