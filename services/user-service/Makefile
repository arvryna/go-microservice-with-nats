protoc:
	mkdir -p pb
	protoc --proto_path=proto proto/*.proto --descriptor_set_out=./pb/descriptor.pb --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb

protob:
   protoc -I${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info \
    --descriptor_set_out=proto.pb proto/userserv.proto