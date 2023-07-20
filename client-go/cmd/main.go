package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hongry18/grpc-example/client/internal/app"
	sampleV1Pb "github.com/hongry18/grpc-example/client/internal/protos/sample-v1"
	sampleV2Pb "github.com/hongry18/grpc-example/client/internal/protos/sample-v2"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	sampleV1Client := sampleV1Pb.NewSampleV1Client(conn)
	sampleV2Client := sampleV2Pb.NewSampleV2Client(conn)
	ctx := context.Background()

	app.GetSampleV1(ctx, sampleV1Client)
	app.GetSampleV2(ctx, sampleV2Client)
}
