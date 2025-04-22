PROTO_DIR=proto
OUT_DIR=.

generate:
	echo "🔧 Generating gRPC code from .proto files..."
	protoc \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/*.proto
	echo "✅ Protobuf generation complete."

clean:
	rm -rf gen/