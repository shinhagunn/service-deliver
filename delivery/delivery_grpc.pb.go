// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: delivery/delivery.proto

package delivery

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

// DeliverClient is the client API for Deliver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliverClient interface {
	Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverResponse, error)
}

type deliverClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliverClient(cc grpc.ClientConnInterface) DeliverClient {
	return &deliverClient{cc}
}

func (c *deliverClient) Deliver(ctx context.Context, in *DeliverRequest, opts ...grpc.CallOption) (*DeliverResponse, error) {
	out := new(DeliverResponse)
	err := c.cc.Invoke(ctx, "/delivery.Deliver/Deliver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliverServer is the server API for Deliver service.
// All implementations must embed UnimplementedDeliverServer
// for forward compatibility
type DeliverServer interface {
	Deliver(context.Context, *DeliverRequest) (*DeliverResponse, error)
	mustEmbedUnimplementedDeliverServer()
}

// UnimplementedDeliverServer must be embedded to have forward compatible implementations.
type UnimplementedDeliverServer struct {
}

func (UnimplementedDeliverServer) Deliver(context.Context, *DeliverRequest) (*DeliverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}
func (UnimplementedDeliverServer) mustEmbedUnimplementedDeliverServer() {}

// UnsafeDeliverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliverServer will
// result in compilation errors.
type UnsafeDeliverServer interface {
	mustEmbedUnimplementedDeliverServer()
}

func RegisterDeliverServer(s grpc.ServiceRegistrar, srv DeliverServer) {
	s.RegisterService(&Deliver_ServiceDesc, srv)
}

func _Deliver_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliverServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delivery.Deliver/Deliver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliverServer).Deliver(ctx, req.(*DeliverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deliver_ServiceDesc is the grpc.ServiceDesc for Deliver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deliver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delivery.Deliver",
	HandlerType: (*DeliverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deliver",
			Handler:    _Deliver_Deliver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery/delivery.proto",
}