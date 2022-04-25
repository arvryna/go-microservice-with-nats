protoc:
	mkdir -p pb
	protoc -I${GOOGLEAPIS_DIR} -I. proto/userserv.proto --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb
