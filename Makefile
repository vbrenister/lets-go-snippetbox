.PHONY: stat_server start_db stop_db

stat_server:
	@echo "Running server..."
	@go run ./cmd/web/

start_db:
	@docker compose up postgres &

stop_db:
	@docker compose down postgres