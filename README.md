# grpc example

## compile
```sh
protoc -I=protos --go_out=server/internal/protos/sample-v1 --go_opt=paths=source_relative \
	    --go-grpc_out=server/internal/protos/sample-v1 --go-grpc_opt=paths=source_relative \
	    sample-v1.proto

protoc -I=protos --go_out=server/internal/protos/sample-v2 --go_opt=paths=source_relative \
	    --go-grpc_out=server/internal/protos/sample-v2 --go-grpc_opt=paths=source_relative \
	    sample-v2.proto
```