.PHONY: help dev-up dev-down dev-clean run dev-fe generate-api build-ui build test

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  %-16s %s\n", $$1, $$2}'

CONFIG ?= .vscode/config.yaml

run: ## Run the Go server (needs dev-up first)
	go run ./cmd/pragmata serve -c $(CONFIG)

dev-fe: ## Run the frontend dev server
	cd frontend && npm run dev

dev-up: ## Start local services (ClickHouse + Zookeeper)
	docker compose -f dev/compose.yaml up -d

dev-down: ## Stop local services
	docker compose -f dev/compose.yaml down

dev-clean: ## Stop local services and delete volumes
	docker compose -f dev/compose.yaml down -v

build-ui: ## Build the React frontend into frontend/dist
	cd frontend && npm ci && npm run build

build: build-ui ## Build the full binary (frontend + Go)
	go build ./cmd/...

test: ## Run Go tests
	go test ./...

generate-api: ## Generate OpenAPI spec and TypeScript API client
	go run ./cmd/pragmata generate openapi
	cd frontend && npm run generate:api
