// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: router.proto

package message

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Acceptor_Routing_FullMethodName = "/router.Acceptor/routing"
)

// AcceptorClient is the client API for Acceptor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcceptorClient interface {
	Routing(ctx context.Context, in *RouteReq, opts ...grpc.CallOption) (*RouteRes, error)
}

type acceptorClient struct {
	cc grpc.ClientConnInterface
}

func NewAcceptorClient(cc grpc.ClientConnInterface) AcceptorClient {
	return &acceptorClient{cc}
}

func (c *acceptorClient) Routing(ctx context.Context, in *RouteReq, opts ...grpc.CallOption) (*RouteRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RouteRes)
	err := c.cc.Invoke(ctx, Acceptor_Routing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcceptorServer is the server API for Acceptor service.
// All implementations must embed UnimplementedAcceptorServer
// for forward compatibility.
type AcceptorServer interface {
	Routing(context.Context, *RouteReq) (*RouteRes, error)
	mustEmbedUnimplementedAcceptorServer()
}

// UnimplementedAcceptorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAcceptorServer struct{}

func (UnimplementedAcceptorServer) Routing(context.Context, *RouteReq) (*RouteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Routing not implemented")
}
func (UnimplementedAcceptorServer) mustEmbedUnimplementedAcceptorServer() {}
func (UnimplementedAcceptorServer) testEmbeddedByValue()                  {}

// UnsafeAcceptorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcceptorServer will
// result in compilation errors.
type UnsafeAcceptorServer interface {
	mustEmbedUnimplementedAcceptorServer()
}

func RegisterAcceptorServer(s grpc.ServiceRegistrar, srv AcceptorServer) {
	// If the following call pancis, it indicates UnimplementedAcceptorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Acceptor_ServiceDesc, srv)
}

func _Acceptor_Routing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcceptorServer).Routing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acceptor_Routing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcceptorServer).Routing(ctx, req.(*RouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Acceptor_ServiceDesc is the grpc.ServiceDesc for Acceptor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Acceptor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "router.Acceptor",
	HandlerType: (*AcceptorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "routing",
			Handler:    _Acceptor_Routing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
