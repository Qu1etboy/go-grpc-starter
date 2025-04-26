PROTO_DIR=proto
OUT_DIR=internal

DB_URL=host=localhost user=test password=test dbname=go-grpc-starter port=5432 sslmode=disable timezone=utc

generate:
	echo "🔧 Generating gRPC code from .proto files..."
	protoc \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/*.proto
	echo "✅ Protobuf generation complete."

clean:
	rm -rf gen/

migrate-up:
	goose -dir ./migrations postgres "$(DB_URL)" up

migrate-status:
	goose -dir ./migrations postgres "$(DB_URL)" status

migrate-down:
	goose -dir ./migrations postgres "$(DB_URL)" down