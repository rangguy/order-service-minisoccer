ifeq ($(OS),Windows_NT)
	GREEN :=
	YELLOW :=
	WHITE :=
	CYAN :=
	RESET :=
else
	GREEN  := $(shell tput -Txterm setaf 2)
	YELLOW := $(shell tput -Txterm setaf 3)
	WHITE  := $(shell tput -Txterm setaf 7)
	CYAN   := $(shell tput -Txterm setaf 6)
	RESET  := $(shell tput -Txterm sgr0)
endif

## Live reload:
watch-prepare: ## Install the tools required for the watch command
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh

watch: ## Run the service with hot reload
	bin/air

## Build:
build: ## Build the service
	go build -o order-service

## Docker:
docker-compose: ## Start the service in docker
	docker-compose up -d --build --force-recreate

docker-build: ## Build the Docker image with a specified tag
	@echo "$(CYAN)Building Docker image...$(RESET)"
	@if [ -z "$(tag)" ]; then \
		echo "$(YELLOW)Error: Please specify the 'tag' parameter, e.g., make docker-build tag=1.0.0$(RESET)"; \
		exit 1; \
	fi
	docker build --platform linux/amd64 -t ranggad/order-service:$(tag) .
	@echo "$(GREEN)Docker image built with tag '$(tag)'$(RESET)"

docker-push: ## Build the Docker image with a specified tag
	@echo "$(CYAN)Building Docker image...$(RESET)"
	@if [ -z "$(tag)" ]; then \
		echo "$(YELLOW)Error: Please specify the 'tag' parameter, e.g., make docker-push tag=1.0.0$(RESET)"; \
		exit 1; \
	fi
	docker push ranggad/order-service:$(tag)
	@echo "$(GREEN)Docker image built with tag '$(tag)'$(RESET)"
