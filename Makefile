.PHONY: dev-up
dev-up:
	@docker compose up -d

.PHONY: dev-down
dev-down:
	@docker compose down

.PHONY: dev-shell
dev-shell:
	@docker compose exec transaction-dev sh

.PHONY: build
build:
	go build -o dist/ ./cmd/...

.PHONY: dev-db-connect
dev-db-connect:
	psql -h postgres -U ${DB_USER} ${DB_NAME}
