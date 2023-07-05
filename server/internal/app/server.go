package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	sampleV1Pb "github.com/hongry18/grpc-example/server/internal/protos/sample-v1"
	sampleV2Pb "github.com/hongry18/grpc-example/server/internal/protos/sample-v2"
)

const defaultPort = "4000"

type serverV1 struct {
	sampleV1Pb.SampleV1Server
}

func (s serverV1) GetSampleV1(ctx context.Context, request *sampleV1Pb.SampleV1Request) (*sampleV1Pb.SampleV1Response, error) {
	return &sampleV1Pb.SampleV1Response{Id: 1, Message: fmt.Sprintf("samplev1, name is %s", request.GetName())}, nil
}

func (s serverV1) GetSampleAgainV1(ctx context.Context, request *sampleV1Pb.SampleV1Request) (*sampleV1Pb.SampleV1Response, error) {
	return &sampleV1Pb.SampleV1Response{Id: 1, Message: fmt.Sprintf("samplev1, name is %s", request.GetName())}, nil
}

type serverV2 struct {
	sampleV2Pb.UnimplementedSampleV2Server
}

func (s serverV2) GetSampleV2(ctx context.Context, request *sampleV2Pb.SampleV2Request) (*sampleV2Pb.SampleV2Response, error) {
	return &sampleV2Pb.SampleV2Response{Id: 1, Message: fmt.Sprintf("samplev2, name is %s", request.GetName())}, nil
}

func (s serverV2) GetSampleAgainV2(ctx context.Context, request *sampleV2Pb.SampleV2Request) (*sampleV2Pb.SampleV2Response, error) {
	return &sampleV2Pb.SampleV2Response{Id: 1, Message: fmt.Sprintf("samplev2, name is %s", request.GetName())}, nil
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
	sampleV1Pb.RegisterSampleV1Server(grpcServer, &serverV1{})
	sampleV2Pb.RegisterSampleV2Server(grpcServer, &serverV2{})

	log.Printf("start gRpc server on %s port", port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
