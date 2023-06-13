compile_grpc:
	protoc --go_out=pkg --go_opt=paths=source_relative \
        --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
     proto/user-service.proto

.PHONY: compile_grpc