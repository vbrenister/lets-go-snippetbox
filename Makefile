.PHONY: start_server start_db stop_db generate_tls test

start_server:
	@echo "Running server..."
	@go run ./cmd/web/

start_db:
	@docker compose up postgres &

stop_db:
	@docker compose down postgres

generate_tls:
	@mkdir tls
	@go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
	@mv cert.pem tls/cert.pem
	@mv key.pem tls/key.pem
	@echo "Certificates generated successfully!"

test:
	@go test -v ./...