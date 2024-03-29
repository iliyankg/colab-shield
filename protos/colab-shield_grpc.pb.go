// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: colab-shield.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ColabShield_HealthCheck_FullMethodName = "/colabshield.ColabShield/HealthCheck"
	ColabShield_Lock_FullMethodName        = "/colabshield.ColabShield/Lock"
)

// ColabShieldClient is the client API for ColabShield service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ColabShieldClient interface {
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
}

type colabShieldClient struct {
	cc grpc.ClientConnInterface
}

func NewColabShieldClient(cc grpc.ClientConnInterface) ColabShieldClient {
	return &colabShieldClient{cc}
}

func (c *colabShieldClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, ColabShield_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colabShieldClient) Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, ColabShield_Lock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ColabShieldServer is the server API for ColabShield service.
// All implementations must embed UnimplementedColabShieldServer
// for forward compatibility
type ColabShieldServer interface {
	HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	Lock(context.Context, *LockRequest) (*LockResponse, error)
	mustEmbedUnimplementedColabShieldServer()
}

// UnimplementedColabShieldServer must be embedded to have forward compatible implementations.
type UnimplementedColabShieldServer struct {
}

func (UnimplementedColabShieldServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedColabShieldServer) Lock(context.Context, *LockRequest) (*LockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedColabShieldServer) mustEmbedUnimplementedColabShieldServer() {}

// UnsafeColabShieldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColabShieldServer will
// result in compilation errors.
type UnsafeColabShieldServer interface {
	mustEmbedUnimplementedColabShieldServer()
}

func RegisterColabShieldServer(s grpc.ServiceRegistrar, srv ColabShieldServer) {
	s.RegisterService(&ColabShield_ServiceDesc, srv)
}

func _ColabShield_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColabShield_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_Lock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).Lock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ColabShield_ServiceDesc is the grpc.ServiceDesc for ColabShield service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColabShield_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colabshield.ColabShield",
	HandlerType: (*ColabShieldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ColabShield_HealthCheck_Handler,
		},
		{
			MethodName: "Lock",
			Handler:    _ColabShield_Lock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "colab-shield.proto",
}