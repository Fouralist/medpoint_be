// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/openapi_merge_a.proto

package examplepb

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

// ServiceAClient is the client API for ServiceA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceAClient interface {
	// ServiceA.MethodOne receives InMessageA and returns OutMessageA
	//
	// Here is the detail explanation about ServiceA.MethodOne.
	MethodOne(ctx context.Context, in *InMessageA, opts ...grpc.CallOption) (*OutMessageA, error)
	// ServiceA.MethodTwo receives OutMessageA and returns InMessageA
	//
	// Here is the detail explanation about ServiceA.MethodTwo.
	MethodTwo(ctx context.Context, in *OutMessageA, opts ...grpc.CallOption) (*InMessageA, error)
}

type serviceAClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAClient(cc grpc.ClientConnInterface) ServiceAClient {
	return &serviceAClient{cc}
}

func (c *serviceAClient) MethodOne(ctx context.Context, in *InMessageA, opts ...grpc.CallOption) (*OutMessageA, error) {
	out := new(OutMessageA)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.ServiceA/MethodOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAClient) MethodTwo(ctx context.Context, in *OutMessageA, opts ...grpc.CallOption) (*InMessageA, error) {
	out := new(InMessageA)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.ServiceA/MethodTwo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAServer is the server API for ServiceA service.
// All implementations should embed UnimplementedServiceAServer
// for forward compatibility
type ServiceAServer interface {
	// ServiceA.MethodOne receives InMessageA and returns OutMessageA
	//
	// Here is the detail explanation about ServiceA.MethodOne.
	MethodOne(context.Context, *InMessageA) (*OutMessageA, error)
	// ServiceA.MethodTwo receives OutMessageA and returns InMessageA
	//
	// Here is the detail explanation about ServiceA.MethodTwo.
	MethodTwo(context.Context, *OutMessageA) (*InMessageA, error)
}

// UnimplementedServiceAServer should be embedded to have forward compatible implementations.
type UnimplementedServiceAServer struct {
}

func (UnimplementedServiceAServer) MethodOne(context.Context, *InMessageA) (*OutMessageA, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodOne not implemented")
}
func (UnimplementedServiceAServer) MethodTwo(context.Context, *OutMessageA) (*InMessageA, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodTwo not implemented")
}

// UnsafeServiceAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceAServer will
// result in compilation errors.
type UnsafeServiceAServer interface {
	mustEmbedUnimplementedServiceAServer()
}

func RegisterServiceAServer(s grpc.ServiceRegistrar, srv ServiceAServer) {
	s.RegisterService(&ServiceA_ServiceDesc, srv)
}

func _ServiceA_MethodOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InMessageA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAServer).MethodOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.ServiceA/MethodOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAServer).MethodOne(ctx, req.(*InMessageA))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceA_MethodTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutMessageA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAServer).MethodTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.ServiceA/MethodTwo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAServer).MethodTwo(ctx, req.(*OutMessageA))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceA_ServiceDesc is the grpc.ServiceDesc for ServiceA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.examplepb.ServiceA",
	HandlerType: (*ServiceAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MethodOne",
			Handler:    _ServiceA_MethodOne_Handler,
		},
		{
			MethodName: "MethodTwo",
			Handler:    _ServiceA_MethodTwo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/openapi_merge_a.proto",
}

// ServiceCClient is the client API for ServiceC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceCClient interface {
	// ServiceC.MethodOne receives InMessageA and returns OutMessageC
	//
	// Here is the detail explanation about ServiceC.MethodOne.
	MethodOne(ctx context.Context, in *InMessageA, opts ...grpc.CallOption) (*OutMessageC, error)
	// ServiceC.MethodTwo receives OutMessageA and returns InMessageA
	//
	// Here is the detail explanation about ServiceC.MethodTwo.
	MethodTwo(ctx context.Context, in *OutMessageA, opts ...grpc.CallOption) (*InMessageA, error)
}

type serviceCClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceCClient(cc grpc.ClientConnInterface) ServiceCClient {
	return &serviceCClient{cc}
}

func (c *serviceCClient) MethodOne(ctx context.Context, in *InMessageA, opts ...grpc.CallOption) (*OutMessageC, error) {
	out := new(OutMessageC)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.ServiceC/MethodOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCClient) MethodTwo(ctx context.Context, in *OutMessageA, opts ...grpc.CallOption) (*InMessageA, error) {
	out := new(InMessageA)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.ServiceC/MethodTwo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceCServer is the server API for ServiceC service.
// All implementations should embed UnimplementedServiceCServer
// for forward compatibility
type ServiceCServer interface {
	// ServiceC.MethodOne receives InMessageA and returns OutMessageC
	//
	// Here is the detail explanation about ServiceC.MethodOne.
	MethodOne(context.Context, *InMessageA) (*OutMessageC, error)
	// ServiceC.MethodTwo receives OutMessageA and returns InMessageA
	//
	// Here is the detail explanation about ServiceC.MethodTwo.
	MethodTwo(context.Context, *OutMessageA) (*InMessageA, error)
}

// UnimplementedServiceCServer should be embedded to have forward compatible implementations.
type UnimplementedServiceCServer struct {
}

func (UnimplementedServiceCServer) MethodOne(context.Context, *InMessageA) (*OutMessageC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodOne not implemented")
}
func (UnimplementedServiceCServer) MethodTwo(context.Context, *OutMessageA) (*InMessageA, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodTwo not implemented")
}

// UnsafeServiceCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceCServer will
// result in compilation errors.
type UnsafeServiceCServer interface {
	mustEmbedUnimplementedServiceCServer()
}

func RegisterServiceCServer(s grpc.ServiceRegistrar, srv ServiceCServer) {
	s.RegisterService(&ServiceC_ServiceDesc, srv)
}

func _ServiceC_MethodOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InMessageA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCServer).MethodOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.ServiceC/MethodOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCServer).MethodOne(ctx, req.(*InMessageA))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceC_MethodTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutMessageA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCServer).MethodTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.ServiceC/MethodTwo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCServer).MethodTwo(ctx, req.(*OutMessageA))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceC_ServiceDesc is the grpc.ServiceDesc for ServiceC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.examplepb.ServiceC",
	HandlerType: (*ServiceCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MethodOne",
			Handler:    _ServiceC_MethodOne_Handler,
		},
		{
			MethodName: "MethodTwo",
			Handler:    _ServiceC_MethodTwo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/openapi_merge_a.proto",
}
