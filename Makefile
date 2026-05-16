.PHONY: dev-up dev-down dev-clean generate-api build-ui build

build-ui: ## Build the React frontend into frontend/dist
	cd frontend && npm ci && npm run build

build: build-ui ## Build the full binary (frontend + Go)
	go build ./cmd/...

generate-api: ## Generate OpenAPI spec and TypeScript API client
	go run ./cmd/pragmata generate openapi
	cd frontend && npm run generate:api

dev-up:
	docker compose -f dev/compose.yaml up -d

dev-down:
	docker compose -f dev/compose.yaml down

dev-clean:
	docker compose -f dev/compose.yaml down -v
