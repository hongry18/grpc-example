package app

import (
	"context"
	"log"
	"time"

	sampleV1 "github.com/hongry18/grpc-example/client/internal/protos/sample-v1"
	sampleV2 "github.com/hongry18/grpc-example/client/internal/protos/sample-v2"
)

func GetSampleV1(ctx context.Context, c sampleV1.SampleV1Client) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	r, err := c.GetSampleV1(ctx, &sampleV1.SampleV1Request{Name: "i am sample v1 client"})
	if err != nil {
		log.Fatalf("could not get sample v1: %v\n", err)
	}
	log.Printf("GetSampleV1: %s", r.GetMessage())
}

func GetSampleV2(ctx context.Context, c sampleV2.SampleV2Client) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	r, err := c.GetSampleV2(ctx, &sampleV2.SampleV2Request{Name: "i am sample v2 client"})
	if err != nil {
		log.Fatalf("could not get sample v2: %v\n", err)
	}
	log.Printf("GetSampleV2: %s", r.GetMessage())
}
