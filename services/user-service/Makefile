# protoc:
# 	mkdir -p pb
# 	protoc --proto_path=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb

protoc:
	mkdir -p pb
	protoc -I${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info --descriptor_set_out=pb/desc.pb proto/userserv.proto --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb
