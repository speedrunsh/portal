// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package portal

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

// PortalClient is the client API for Portal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalClient interface {
	ServiceRestart(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error)
	ServiceStart(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error)
	ServiceStop(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error)
	ServiceStatus(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error)
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
}

type portalClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalClient(cc grpc.ClientConnInterface) PortalClient {
	return &portalClient{cc}
}

func (c *portalClient) ServiceRestart(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceRestart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) ServiceStart(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) ServiceStop(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) ServiceStatus(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portal.Portal/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portal.Portal/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalServer is the server API for Portal service.
// All implementations must embed UnimplementedPortalServer
// for forward compatibility
type PortalServer interface {
	ServiceRestart(context.Context, *Service) (*Response, error)
	ServiceStart(context.Context, *Service) (*Response, error)
	ServiceStop(context.Context, *Service) (*Response, error)
	ServiceStatus(context.Context, *Service) (*Response, error)
	RunCommand(context.Context, *Command) (*Response, error)
	mustEmbedUnimplementedPortalServer()
}

// UnimplementedPortalServer must be embedded to have forward compatible implementations.
type UnimplementedPortalServer struct {
}

func (UnimplementedPortalServer) ServiceRestart(context.Context, *Service) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceRestart not implemented")
}
func (UnimplementedPortalServer) ServiceStart(context.Context, *Service) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStart not implemented")
}
func (UnimplementedPortalServer) ServiceStop(context.Context, *Service) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStop not implemented")
}
func (UnimplementedPortalServer) ServiceStatus(context.Context, *Service) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}
func (UnimplementedPortalServer) RunCommand(context.Context, *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (UnimplementedPortalServer) mustEmbedUnimplementedPortalServer() {}

// UnsafePortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalServer will
// result in compilation errors.
type UnsafePortalServer interface {
	mustEmbedUnimplementedPortalServer()
}

func RegisterPortalServer(s grpc.ServiceRegistrar, srv PortalServer) {
	s.RegisterService(&Portal_ServiceDesc, srv)
}

func _Portal_ServiceRestart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).ServiceRestart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/ServiceRestart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).ServiceRestart(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_ServiceStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).ServiceStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/ServiceStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).ServiceStart(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_ServiceStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).ServiceStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/ServiceStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).ServiceStop(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).ServiceStatus(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portal.Portal/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

// Portal_ServiceDesc is the grpc.ServiceDesc for Portal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Portal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "portal.Portal",
	HandlerType: (*PortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceRestart",
			Handler:    _Portal_ServiceRestart_Handler,
		},
		{
			MethodName: "ServiceStart",
			Handler:    _Portal_ServiceStart_Handler,
		},
		{
			MethodName: "ServiceStop",
			Handler:    _Portal_ServiceStop_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Portal_ServiceStatus_Handler,
		},
		{
			MethodName: "RunCommand",
			Handler:    _Portal_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal.proto",
}
