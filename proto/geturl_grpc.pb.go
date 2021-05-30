// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SendAddressClient is the client API for SendAddress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendAddressClient interface {
	GetUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlReply, error)
}

type sendAddressClient struct {
	cc grpc.ClientConnInterface
}

func NewSendAddressClient(cc grpc.ClientConnInterface) SendAddressClient {
	return &sendAddressClient{cc}
}

func (c *sendAddressClient) GetUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlReply, error) {
	out := new(UrlReply)
	err := c.cc.Invoke(ctx, "/proto.SendAddress/GetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendAddressServer is the server API for SendAddress service.
// All implementations must embed UnimplementedSendAddressServer
// for forward compatibility
type SendAddressServer interface {
	GetUrl(context.Context, *UrlRequest) (*UrlReply, error)
	mustEmbedUnimplementedSendAddressServer()
}

// UnimplementedSendAddressServer must be embedded to have forward compatible implementations.
type UnimplementedSendAddressServer struct {
}

func (UnimplementedSendAddressServer) GetUrl(context.Context, *UrlRequest) (*UrlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedSendAddressServer) mustEmbedUnimplementedSendAddressServer() {}

// UnsafeSendAddressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendAddressServer will
// result in compilation errors.
type UnsafeSendAddressServer interface {
	mustEmbedUnimplementedSendAddressServer()
}

func RegisterSendAddressServer(s grpc.ServiceRegistrar, srv SendAddressServer) {
	s.RegisterService(&SendAddress_ServiceDesc, srv)
}

func _SendAddress_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendAddressServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SendAddress/GetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendAddressServer).GetUrl(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SendAddress_ServiceDesc is the grpc.ServiceDesc for SendAddress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendAddress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SendAddress",
	HandlerType: (*SendAddressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUrl",
			Handler:    _SendAddress_GetUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geturl.proto",
}
