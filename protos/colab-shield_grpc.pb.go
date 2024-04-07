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
	ColabShield_HealthCheck_FullMethodName  = "/colabshield.ColabShield/HealthCheck"
	ColabShield_InitProject_FullMethodName  = "/colabshield.ColabShield/InitProject"
	ColabShield_ListProjects_FullMethodName = "/colabshield.ColabShield/ListProjects"
	ColabShield_ListFiles_FullMethodName    = "/colabshield.ColabShield/ListFiles"
	ColabShield_Claim_FullMethodName        = "/colabshield.ColabShield/Claim"
	ColabShield_Update_FullMethodName       = "/colabshield.ColabShield/Update"
	ColabShield_Release_FullMethodName      = "/colabshield.ColabShield/Release"
)

// ColabShieldClient is the client API for ColabShield service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ColabShieldClient interface {
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	InitProject(ctx context.Context, in *InitProjectRequest, opts ...grpc.CallOption) (*InitProjectResponse, error)
	ListProjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	Claim(ctx context.Context, in *ClaimFilesRequest, opts ...grpc.CallOption) (*ClaimFilesResponse, error)
	Update(ctx context.Context, in *UpdateFilesRequest, opts ...grpc.CallOption) (*UpdateFilesResponse, error)
	Release(ctx context.Context, in *ReleaseFilesRequest, opts ...grpc.CallOption) (*ReleaseFilesResponse, error)
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

func (c *colabShieldClient) InitProject(ctx context.Context, in *InitProjectRequest, opts ...grpc.CallOption) (*InitProjectResponse, error) {
	out := new(InitProjectResponse)
	err := c.cc.Invoke(ctx, ColabShield_InitProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colabShieldClient) ListProjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, ColabShield_ListProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colabShieldClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, ColabShield_ListFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colabShieldClient) Claim(ctx context.Context, in *ClaimFilesRequest, opts ...grpc.CallOption) (*ClaimFilesResponse, error) {
	out := new(ClaimFilesResponse)
	err := c.cc.Invoke(ctx, ColabShield_Claim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colabShieldClient) Update(ctx context.Context, in *UpdateFilesRequest, opts ...grpc.CallOption) (*UpdateFilesResponse, error) {
	out := new(UpdateFilesResponse)
	err := c.cc.Invoke(ctx, ColabShield_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colabShieldClient) Release(ctx context.Context, in *ReleaseFilesRequest, opts ...grpc.CallOption) (*ReleaseFilesResponse, error) {
	out := new(ReleaseFilesResponse)
	err := c.cc.Invoke(ctx, ColabShield_Release_FullMethodName, in, out, opts...)
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
	InitProject(context.Context, *InitProjectRequest) (*InitProjectResponse, error)
	ListProjects(context.Context, *emptypb.Empty) (*ListProjectsResponse, error)
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	Claim(context.Context, *ClaimFilesRequest) (*ClaimFilesResponse, error)
	Update(context.Context, *UpdateFilesRequest) (*UpdateFilesResponse, error)
	Release(context.Context, *ReleaseFilesRequest) (*ReleaseFilesResponse, error)
	mustEmbedUnimplementedColabShieldServer()
}

// UnimplementedColabShieldServer must be embedded to have forward compatible implementations.
type UnimplementedColabShieldServer struct {
}

func (UnimplementedColabShieldServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedColabShieldServer) InitProject(context.Context, *InitProjectRequest) (*InitProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitProject not implemented")
}
func (UnimplementedColabShieldServer) ListProjects(context.Context, *emptypb.Empty) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedColabShieldServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedColabShieldServer) Claim(context.Context, *ClaimFilesRequest) (*ClaimFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
}
func (UnimplementedColabShieldServer) Update(context.Context, *UpdateFilesRequest) (*UpdateFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedColabShieldServer) Release(context.Context, *ReleaseFilesRequest) (*ReleaseFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
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

func _ColabShield_InitProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).InitProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_InitProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).InitProject(ctx, req.(*InitProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColabShield_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).ListProjects(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColabShield_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_ListFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColabShield_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_Claim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).Claim(ctx, req.(*ClaimFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColabShield_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).Update(ctx, req.(*UpdateFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColabShield_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColabShieldServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColabShield_Release_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColabShieldServer).Release(ctx, req.(*ReleaseFilesRequest))
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
			MethodName: "InitProject",
			Handler:    _ColabShield_InitProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _ColabShield_ListProjects_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _ColabShield_ListFiles_Handler,
		},
		{
			MethodName: "Claim",
			Handler:    _ColabShield_Claim_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ColabShield_Update_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _ColabShield_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "colab-shield.proto",
}
