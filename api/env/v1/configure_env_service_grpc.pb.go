// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: configure_env_service.proto

package v1

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
	Service_AllEnv_FullMethodName        = "/configure_env.Service/AllEnv"
	Service_AddEnv_FullMethodName        = "/configure_env.Service/AddEnv"
	Service_UpdateEnv_FullMethodName     = "/configure_env.Service/UpdateEnv"
	Service_DeleteEnv_FullMethodName     = "/configure_env.Service/DeleteEnv"
	Service_GetEnvToken_FullMethodName   = "/configure_env.Service/GetEnvToken"
	Service_ResetEnvToken_FullMethodName = "/configure_env.Service/ResetEnvToken"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// AllEnv 获取全部环境
	AllEnv(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllEnvReply, error)
	// AddEnv 添加环境
	AddEnv(ctx context.Context, in *AddEnvRequest, opts ...grpc.CallOption) (*AddEnvReply, error)
	// UpdateEnv 更新环境信息
	UpdateEnv(ctx context.Context, in *UpdateEnvRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteEnv 删除环境信息
	DeleteEnv(ctx context.Context, in *DeleteEnvRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetEnvToken 获取环境token
	GetEnvToken(ctx context.Context, in *GetEnvTokenRequest, opts ...grpc.CallOption) (*GetEnvTokenReply, error)
	// ResetEnvToken 重置环境token
	ResetEnvToken(ctx context.Context, in *ResetEnvTokenRequest, opts ...grpc.CallOption) (*ResetEnvTokenReply, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AllEnv(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllEnvReply, error) {
	out := new(AllEnvReply)
	err := c.cc.Invoke(ctx, Service_AllEnv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddEnv(ctx context.Context, in *AddEnvRequest, opts ...grpc.CallOption) (*AddEnvReply, error) {
	out := new(AddEnvReply)
	err := c.cc.Invoke(ctx, Service_AddEnv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateEnv(ctx context.Context, in *UpdateEnvRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateEnv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteEnv(ctx context.Context, in *DeleteEnvRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteEnv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetEnvToken(ctx context.Context, in *GetEnvTokenRequest, opts ...grpc.CallOption) (*GetEnvTokenReply, error) {
	out := new(GetEnvTokenReply)
	err := c.cc.Invoke(ctx, Service_GetEnvToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ResetEnvToken(ctx context.Context, in *ResetEnvTokenRequest, opts ...grpc.CallOption) (*ResetEnvTokenReply, error) {
	out := new(ResetEnvTokenReply)
	err := c.cc.Invoke(ctx, Service_ResetEnvToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// AllEnv 获取全部环境
	AllEnv(context.Context, *emptypb.Empty) (*AllEnvReply, error)
	// AddEnv 添加环境
	AddEnv(context.Context, *AddEnvRequest) (*AddEnvReply, error)
	// UpdateEnv 更新环境信息
	UpdateEnv(context.Context, *UpdateEnvRequest) (*emptypb.Empty, error)
	// DeleteEnv 删除环境信息
	DeleteEnv(context.Context, *DeleteEnvRequest) (*emptypb.Empty, error)
	// GetEnvToken 获取环境token
	GetEnvToken(context.Context, *GetEnvTokenRequest) (*GetEnvTokenReply, error)
	// ResetEnvToken 重置环境token
	ResetEnvToken(context.Context, *ResetEnvTokenRequest) (*ResetEnvTokenReply, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) AllEnv(context.Context, *emptypb.Empty) (*AllEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllEnv not implemented")
}
func (UnimplementedServiceServer) AddEnv(context.Context, *AddEnvRequest) (*AddEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEnv not implemented")
}
func (UnimplementedServiceServer) UpdateEnv(context.Context, *UpdateEnvRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnv not implemented")
}
func (UnimplementedServiceServer) DeleteEnv(context.Context, *DeleteEnvRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEnv not implemented")
}
func (UnimplementedServiceServer) GetEnvToken(context.Context, *GetEnvTokenRequest) (*GetEnvTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvToken not implemented")
}
func (UnimplementedServiceServer) ResetEnvToken(context.Context, *ResetEnvTokenRequest) (*ResetEnvTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetEnvToken not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_AllEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AllEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AllEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AllEnv(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddEnv(ctx, req.(*AddEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateEnv(ctx, req.(*UpdateEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteEnv(ctx, req.(*DeleteEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetEnvToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetEnvToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetEnvToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetEnvToken(ctx, req.(*GetEnvTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ResetEnvToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetEnvTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ResetEnvToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ResetEnvToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ResetEnvToken(ctx, req.(*ResetEnvTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configure_env.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllEnv",
			Handler:    _Service_AllEnv_Handler,
		},
		{
			MethodName: "AddEnv",
			Handler:    _Service_AddEnv_Handler,
		},
		{
			MethodName: "UpdateEnv",
			Handler:    _Service_UpdateEnv_Handler,
		},
		{
			MethodName: "DeleteEnv",
			Handler:    _Service_DeleteEnv_Handler,
		},
		{
			MethodName: "GetEnvToken",
			Handler:    _Service_GetEnvToken_Handler,
		},
		{
			MethodName: "ResetEnvToken",
			Handler:    _Service_ResetEnvToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configure_env_service.proto",
}
