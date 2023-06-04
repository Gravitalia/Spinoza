// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: proto/spinoza.proto

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

// SpinozaClient is the client API for Spinoza service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpinozaClient interface {
	// Compress and then upload the image to the CDN provider
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadReply, error)
}

type spinozaClient struct {
	cc grpc.ClientConnInterface
}

func NewSpinozaClient(cc grpc.ClientConnInterface) SpinozaClient {
	return &spinozaClient{cc}
}

func (c *spinozaClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadReply, error) {
	out := new(UploadReply)
	err := c.cc.Invoke(ctx, "/spinoza.Spinoza/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpinozaServer is the server API for Spinoza service.
// All implementations must embed UnimplementedSpinozaServer
// for forward compatibility
type SpinozaServer interface {
	// Compress and then upload the image to the CDN provider
	Upload(context.Context, *UploadRequest) (*UploadReply, error)
	mustEmbedUnimplementedSpinozaServer()
}

// UnimplementedSpinozaServer must be embedded to have forward compatible implementations.
type UnimplementedSpinozaServer struct {
}

func (UnimplementedSpinozaServer) Upload(context.Context, *UploadRequest) (*UploadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedSpinozaServer) mustEmbedUnimplementedSpinozaServer() {}

// UnsafeSpinozaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpinozaServer will
// result in compilation errors.
type UnsafeSpinozaServer interface {
	mustEmbedUnimplementedSpinozaServer()
}

func RegisterSpinozaServer(s grpc.ServiceRegistrar, srv SpinozaServer) {
	s.RegisterService(&Spinoza_ServiceDesc, srv)
}

func _Spinoza_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpinozaServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spinoza.Spinoza/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpinozaServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Spinoza_ServiceDesc is the grpc.ServiceDesc for Spinoza service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Spinoza_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spinoza.Spinoza",
	HandlerType: (*SpinozaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _Spinoza_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/spinoza.proto",
}
