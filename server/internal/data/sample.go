package data

import (
	samplePb "github.com/hongry18/grpc-example/internal/protos/sample"
)

var sampleData = []*samplePb.SampleResponse{
	{Id: 1, Message: "msg-1"},
	{Id: 2, Message: "msg-2"},
	{Id: 3, Message: "msg-3"},
	{Id: 4, Message: "msg-4"},
}
