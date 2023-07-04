package app

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	sampleV1Pb "github.com/hongry18/grpc-example/server/internal/protos/sample-v1"
	sampleV2Pb "github.com/hongry18/grpc-example/server/internal/protos/sample-v2"
)

const defaultPort = "9000"

type serverV1 struct {
	sampleV1Pb.UnimplementedSampleV1StoreServer
}

type serverV2 struct {
	sampleV2Pb.UnimplementedSampleV2StoreServer
}

func StartServer(port string) {
	if port == "" {
		port = defaultPort
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	sampleV1Pb.RegisterSampleV1StoreServer(grpcServer, &serverV1{})
	sampleV2Pb.RegisterSampleV2StoreServer(grpcServer, &serverV2{})

	log.Printf("start gRpc server on %s port", port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
