gen-orders:
	@protoc \
	--proto_path=./protobuf "./protobuf/orders.proto" \
	--go_out=./services/common/genproto/orders --go_opt=paths=source_relative \
	--go-grpc_out=./services/common/genproto/orders --go-grpc_opt=paths=source_relative
run-orders:
	@go run ./services/orders

gen-hello:
	@protoc \
	--proto_path=./protobuf "./protobuf/hello.proto" \
	--go_out=./services/common/genproto/hello --go_opt=paths=source_relative \
	--go-grpc_out=./services/common/genproto/hello --go-grpc_opt=paths=source_relative
run-hello:
	@go run ./services/hello