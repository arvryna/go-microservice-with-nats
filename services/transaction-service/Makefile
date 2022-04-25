protoc:
	mkdir -p pb
	protoc -I${GOOGLEAPIS_DIR} -I. proto/transervice.proto --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb
