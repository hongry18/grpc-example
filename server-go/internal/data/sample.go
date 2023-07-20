package data

import (
	samplePbV1 "github.com/hongry18/grpc-example/server/internal/protos/sample-v1"
)

var sampleData = []*samplePbV1.SampleV1Response{
	{Id: 1, Message: "msg-1"},
	{Id: 2, Message: "msg-2"},
	{Id: 3, Message: "msg-3"},
	{Id: 4, Message: "msg-4"},
}
