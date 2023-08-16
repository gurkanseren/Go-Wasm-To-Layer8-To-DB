// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/Layer8Slave.proto

package service

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

// Layer8MasterServiceClient is the client API for Layer8MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Layer8MasterServiceClient interface {
	GetJwtSecret(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*JwtSecretResponse, error)
	GetPublicKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PublicKeyResponse, error)
}

type layer8MasterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLayer8MasterServiceClient(cc grpc.ClientConnInterface) Layer8MasterServiceClient {
	return &layer8MasterServiceClient{cc}
}

func (c *layer8MasterServiceClient) GetJwtSecret(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*JwtSecretResponse, error) {
	out := new(JwtSecretResponse)
	err := c.cc.Invoke(ctx, "/Layer8MasterService/GetJwtSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layer8MasterServiceClient) GetPublicKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PublicKeyResponse, error) {
	out := new(PublicKeyResponse)
	err := c.cc.Invoke(ctx, "/Layer8MasterService/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Layer8MasterServiceServer is the server API for Layer8MasterService service.
// All implementations must embed UnimplementedLayer8MasterServiceServer
// for forward compatibility
type Layer8MasterServiceServer interface {
	GetJwtSecret(context.Context, *Empty) (*JwtSecretResponse, error)
	GetPublicKey(context.Context, *Empty) (*PublicKeyResponse, error)
	mustEmbedUnimplementedLayer8MasterServiceServer()
}

// UnimplementedLayer8MasterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLayer8MasterServiceServer struct {
}

func (UnimplementedLayer8MasterServiceServer) GetJwtSecret(context.Context, *Empty) (*JwtSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJwtSecret not implemented")
}
func (UnimplementedLayer8MasterServiceServer) GetPublicKey(context.Context, *Empty) (*PublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedLayer8MasterServiceServer) mustEmbedUnimplementedLayer8MasterServiceServer() {}

// UnsafeLayer8MasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Layer8MasterServiceServer will
// result in compilation errors.
type UnsafeLayer8MasterServiceServer interface {
	mustEmbedUnimplementedLayer8MasterServiceServer()
}

func RegisterLayer8MasterServiceServer(s grpc.ServiceRegistrar, srv Layer8MasterServiceServer) {
	s.RegisterService(&Layer8MasterService_ServiceDesc, srv)
}

func _Layer8MasterService_GetJwtSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Layer8MasterServiceServer).GetJwtSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Layer8MasterService/GetJwtSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Layer8MasterServiceServer).GetJwtSecret(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layer8MasterService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Layer8MasterServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Layer8MasterService/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Layer8MasterServiceServer).GetPublicKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Layer8MasterService_ServiceDesc is the grpc.ServiceDesc for Layer8MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Layer8MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Layer8MasterService",
	HandlerType: (*Layer8MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJwtSecret",
			Handler:    _Layer8MasterService_GetJwtSecret_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _Layer8MasterService_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Layer8Slave.proto",
}
