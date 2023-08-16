// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/kv_service.proto

package proto

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

// KvServiceClient is the client API for KvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KvServiceClient interface {
	GetVal(ctx context.Context, in *GetKvRequest, opts ...grpc.CallOption) (*GetKvResponse, error)
	SetVal(ctx context.Context, in *SetKvRequest, opts ...grpc.CallOption) (*SetKvResponse, error)
}

type kvServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKvServiceClient(cc grpc.ClientConnInterface) KvServiceClient {
	return &kvServiceClient{cc}
}

func (c *kvServiceClient) GetVal(ctx context.Context, in *GetKvRequest, opts ...grpc.CallOption) (*GetKvResponse, error) {
	out := new(GetKvResponse)
	err := c.cc.Invoke(ctx, "/kv_service.KvService/GetVal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvServiceClient) SetVal(ctx context.Context, in *SetKvRequest, opts ...grpc.CallOption) (*SetKvResponse, error) {
	out := new(SetKvResponse)
	err := c.cc.Invoke(ctx, "/kv_service.KvService/SetVal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KvServiceServer is the server API for KvService service.
// All implementations must embed UnimplementedKvServiceServer
// for forward compatibility
type KvServiceServer interface {
	GetVal(context.Context, *GetKvRequest) (*GetKvResponse, error)
	SetVal(context.Context, *SetKvRequest) (*SetKvResponse, error)
	mustEmbedUnimplementedKvServiceServer()
}

// UnimplementedKvServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKvServiceServer struct {
}

func (UnimplementedKvServiceServer) GetVal(context.Context, *GetKvRequest) (*GetKvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVal not implemented")
}
func (UnimplementedKvServiceServer) SetVal(context.Context, *SetKvRequest) (*SetKvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVal not implemented")
}
func (UnimplementedKvServiceServer) mustEmbedUnimplementedKvServiceServer() {}

// UnsafeKvServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KvServiceServer will
// result in compilation errors.
type UnsafeKvServiceServer interface {
	mustEmbedUnimplementedKvServiceServer()
}

func RegisterKvServiceServer(s grpc.ServiceRegistrar, srv KvServiceServer) {
	s.RegisterService(&KvService_ServiceDesc, srv)
}

func _KvService_GetVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServiceServer).GetVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv_service.KvService/GetVal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServiceServer).GetVal(ctx, req.(*GetKvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KvService_SetVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServiceServer).SetVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv_service.KvService/SetVal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServiceServer).SetVal(ctx, req.(*SetKvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KvService_ServiceDesc is the grpc.ServiceDesc for KvService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KvService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kv_service.KvService",
	HandlerType: (*KvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVal",
			Handler:    _KvService_GetVal_Handler,
		},
		{
			MethodName: "SetVal",
			Handler:    _KvService_SetVal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/kv_service.proto",
}
