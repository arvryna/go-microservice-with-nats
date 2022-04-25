protoc:
	protoc -I${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info --descriptor_set_out=desc.pb services/user-service/proto/userserv.proto services/transaction-service/proto/transervice.proto
