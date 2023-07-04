# grpc example

## compile
```sh
# server
protoc -I=protos \
	--go_out=server/internal/protos \
	--go_opt=Msample-v1/sample-v1.proto=./sample-v1 \
	--go-grpc_out=server/internal/protos \
	--go-grpc_opt=Msample-v1/sample-v1.proto=./sample-v1 \
	sample-v1/sample-v1.proto

protoc -I=protos \
	--go_out=server/internal/protos \
	--go_opt=Msample-v2/sample-v2.proto=./sample-v2 \
	--go-grpc_out=server/internal/protos \
	--go-grpc_opt=Msample-v2/sample-v2.proto=./sample-v2 \
	sample-v2/sample-v2.proto

# client
protoc -I=protos \
	--go_out=client/internal/protos \
	--go_opt=Msample-v1/sample-v1.proto=./sample-v1 \
	--go-grpc_out=client/internal/protos \
	--go-grpc_opt=Msample-v1/sample-v1.proto=./sample-v1 \
	sample-v1/sample-v1.proto

protoc -I=protos \
	--go_out=client/internal/protos \
	--go_opt=Msample-v2/sample-v2.proto=./sample-v2 \
	--go-grpc_out=client/internal/protos \
	--go-grpc_opt=Msample-v2/sample-v2.proto=./sample-v2 \
	sample-v2/sample-v2.proto
```