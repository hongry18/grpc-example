// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: sample-v1/sample-v1.proto

package sample_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SampleV1Store_GetSampleV1_FullMethodName      = "/sample_v1.SampleV1Store/GetSampleV1"
	SampleV1Store_GetSampleAgainV1_FullMethodName = "/sample_v1.SampleV1Store/GetSampleAgainV1"
)

// SampleV1StoreClient is the client API for SampleV1Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleV1StoreClient interface {
	GetSampleV1(ctx context.Context, in *SampleV1Request, opts ...grpc.CallOption) (*SampleV1Response, error)
	GetSampleAgainV1(ctx context.Context, in *SampleV1Request, opts ...grpc.CallOption) (*SampleV1Response, error)
}

type sampleV1StoreClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleV1StoreClient(cc grpc.ClientConnInterface) SampleV1StoreClient {
	return &sampleV1StoreClient{cc}
}

func (c *sampleV1StoreClient) GetSampleV1(ctx context.Context, in *SampleV1Request, opts ...grpc.CallOption) (*SampleV1Response, error) {
	out := new(SampleV1Response)
	err := c.cc.Invoke(ctx, SampleV1Store_GetSampleV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleV1StoreClient) GetSampleAgainV1(ctx context.Context, in *SampleV1Request, opts ...grpc.CallOption) (*SampleV1Response, error) {
	out := new(SampleV1Response)
	err := c.cc.Invoke(ctx, SampleV1Store_GetSampleAgainV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleV1StoreServer is the server API for SampleV1Store service.
// All implementations must embed UnimplementedSampleV1StoreServer
// for forward compatibility
type SampleV1StoreServer interface {
	GetSampleV1(context.Context, *SampleV1Request) (*SampleV1Response, error)
	GetSampleAgainV1(context.Context, *SampleV1Request) (*SampleV1Response, error)
	mustEmbedUnimplementedSampleV1StoreServer()
}

// UnimplementedSampleV1StoreServer must be embedded to have forward compatible implementations.
type UnimplementedSampleV1StoreServer struct {
}

func (UnimplementedSampleV1StoreServer) GetSampleV1(context.Context, *SampleV1Request) (*SampleV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSampleV1 not implemented")
}
func (UnimplementedSampleV1StoreServer) GetSampleAgainV1(context.Context, *SampleV1Request) (*SampleV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSampleAgainV1 not implemented")
}
func (UnimplementedSampleV1StoreServer) mustEmbedUnimplementedSampleV1StoreServer() {}

// UnsafeSampleV1StoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleV1StoreServer will
// result in compilation errors.
type UnsafeSampleV1StoreServer interface {
	mustEmbedUnimplementedSampleV1StoreServer()
}

func RegisterSampleV1StoreServer(s grpc.ServiceRegistrar, srv SampleV1StoreServer) {
	s.RegisterService(&SampleV1Store_ServiceDesc, srv)
}

func _SampleV1Store_GetSampleV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleV1StoreServer).GetSampleV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SampleV1Store_GetSampleV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleV1StoreServer).GetSampleV1(ctx, req.(*SampleV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleV1Store_GetSampleAgainV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleV1StoreServer).GetSampleAgainV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SampleV1Store_GetSampleAgainV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleV1StoreServer).GetSampleAgainV1(ctx, req.(*SampleV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleV1Store_ServiceDesc is the grpc.ServiceDesc for SampleV1Store service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleV1Store_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample_v1.SampleV1Store",
	HandlerType: (*SampleV1StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSampleV1",
			Handler:    _SampleV1Store_GetSampleV1_Handler,
		},
		{
			MethodName: "GetSampleAgainV1",
			Handler:    _SampleV1Store_GetSampleAgainV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample-v1/sample-v1.proto",
}
