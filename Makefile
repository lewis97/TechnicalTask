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
