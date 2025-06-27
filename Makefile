gen-orders:
	@protoc \
	--proto_path=./protobuf "./protobuf/orders.proto" \
	--go_out=./services/common/protobuf/orders/ --go_opt=paths=source_relative \
	--go-grpc_out=./services/common/protobuf/orders/ --go-grpc_opt=paths=source_relative
gen-hello:
	@protoc \
	--proto_path=./protobuf "./protobuf/hello.proto" \
	--go_out=./services/common/protobuf/hello/ --go_opt=paths=source_relative \
	--go-grpc_out=./services/common/protobuf/hello/ --go-grpc_opt=paths=source_relative