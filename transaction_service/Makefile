compile_grpc:
	protoc --go_out=pkg --go_opt=paths=source_relative \
        --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
     proto/transaction_service.proto

.PHONY: compile_grpc