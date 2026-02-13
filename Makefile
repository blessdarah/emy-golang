.PHONY: help dev prod down logs rebuild clean ps

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Start development environment with hot reload
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build

prod: ## Start production environment
	docker-compose up -d --build

down: ## Stop and remove all containers
	docker-compose down

logs: ## Show logs from all containers
	docker-compose logs -f

logs-app: ## Show logs from app container only
	docker logs -f go_book_api

rebuild: down dev ## Rebuild and restart development environment

clean: ## Remove all containers, volumes, and images
	docker-compose down -v --rmi all

ps: ## Show running containers
	docker-compose ps

watch: dev logs-app ## Start dev mode and watch app logs
