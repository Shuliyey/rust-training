// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: resize.proto

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

// ResizeClient is the client API for Resize service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResizeClient interface {
	// function which can be called
	Send(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*ResizeResponse, error)
}

type resizeClient struct {
	cc grpc.ClientConnInterface
}

func NewResizeClient(cc grpc.ClientConnInterface) ResizeClient {
	return &resizeClient{cc}
}

func (c *resizeClient) Send(ctx context.Context, in *ResizeRequest, opts ...grpc.CallOption) (*ResizeResponse, error) {
	out := new(ResizeResponse)
	err := c.cc.Invoke(ctx, "/resize.Resize/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResizeServer is the server API for Resize service.
// All implementations must embed UnimplementedResizeServer
// for forward compatibility
type ResizeServer interface {
	// function which can be called
	Send(context.Context, *ResizeRequest) (*ResizeResponse, error)
	mustEmbedUnimplementedResizeServer()
}

// UnimplementedResizeServer must be embedded to have forward compatible implementations.
type UnimplementedResizeServer struct {
}

func (UnimplementedResizeServer) Send(context.Context, *ResizeRequest) (*ResizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedResizeServer) mustEmbedUnimplementedResizeServer() {}

// UnsafeResizeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResizeServer will
// result in compilation errors.
type UnsafeResizeServer interface {
	mustEmbedUnimplementedResizeServer()
}

func RegisterResizeServer(s grpc.ServiceRegistrar, srv ResizeServer) {
	s.RegisterService(&Resize_ServiceDesc, srv)
}

func _Resize_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResizeServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resize.Resize/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResizeServer).Send(ctx, req.(*ResizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resize_ServiceDesc is the grpc.ServiceDesc for Resize service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resize_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resize.Resize",
	HandlerType: (*ResizeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Resize_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resize.proto",
}