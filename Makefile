.PHONY: dev-up dev-down dev-clean

dev-up:
	docker compose -f dev/compose.yaml up -d

dev-down:
	docker compose -f dev/compose.yaml down

dev-clean:
	docker compose -f dev/compose.yaml down -v
