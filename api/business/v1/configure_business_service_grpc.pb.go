// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: configure_business_service.proto

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
	Service_PageBusiness_FullMethodName         = "/configure_business.Service/PageBusiness"
	Service_AddBusiness_FullMethodName          = "/configure_business.Service/AddBusiness"
	Service_UpdateBusiness_FullMethodName       = "/configure_business.Service/UpdateBusiness"
	Service_DeleteBusiness_FullMethodName       = "/configure_business.Service/DeleteBusiness"
	Service_GetBusinessValues_FullMethodName    = "/configure_business.Service/GetBusinessValues"
	Service_UpdateBusinessValues_FullMethodName = "/configure_business.Service/UpdateBusinessValues"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	PageBusiness(ctx context.Context, in *PageBusinessRequest, opts ...grpc.CallOption) (*PageBusinessReply, error)
	AddBusiness(ctx context.Context, in *AddBusinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBusiness(ctx context.Context, in *UpdateBusinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteBusiness(ctx context.Context, in *DeleteBusinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBusinessValues(ctx context.Context, in *GetBusinessValuesRequest, opts ...grpc.CallOption) (*GetBusinessValuesReply, error)
	UpdateBusinessValues(ctx context.Context, in *UpdateBusinessValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) PageBusiness(ctx context.Context, in *PageBusinessRequest, opts ...grpc.CallOption) (*PageBusinessReply, error) {
	out := new(PageBusinessReply)
	err := c.cc.Invoke(ctx, Service_PageBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddBusiness(ctx context.Context, in *AddBusinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_AddBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateBusiness(ctx context.Context, in *UpdateBusinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteBusiness(ctx context.Context, in *DeleteBusinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetBusinessValues(ctx context.Context, in *GetBusinessValuesRequest, opts ...grpc.CallOption) (*GetBusinessValuesReply, error) {
	out := new(GetBusinessValuesReply)
	err := c.cc.Invoke(ctx, Service_GetBusinessValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateBusinessValues(ctx context.Context, in *UpdateBusinessValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateBusinessValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	PageBusiness(context.Context, *PageBusinessRequest) (*PageBusinessReply, error)
	AddBusiness(context.Context, *AddBusinessRequest) (*emptypb.Empty, error)
	UpdateBusiness(context.Context, *UpdateBusinessRequest) (*emptypb.Empty, error)
	DeleteBusiness(context.Context, *DeleteBusinessRequest) (*emptypb.Empty, error)
	GetBusinessValues(context.Context, *GetBusinessValuesRequest) (*GetBusinessValuesReply, error)
	UpdateBusinessValues(context.Context, *UpdateBusinessValueRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) PageBusiness(context.Context, *PageBusinessRequest) (*PageBusinessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageBusiness not implemented")
}
func (UnimplementedServiceServer) AddBusiness(context.Context, *AddBusinessRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBusiness not implemented")
}
func (UnimplementedServiceServer) UpdateBusiness(context.Context, *UpdateBusinessRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusiness not implemented")
}
func (UnimplementedServiceServer) DeleteBusiness(context.Context, *DeleteBusinessRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBusiness not implemented")
}
func (UnimplementedServiceServer) GetBusinessValues(context.Context, *GetBusinessValuesRequest) (*GetBusinessValuesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessValues not implemented")
}
func (UnimplementedServiceServer) UpdateBusinessValues(context.Context, *UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusinessValues not implemented")
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

func _Service_PageBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PageBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_PageBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PageBusiness(ctx, req.(*PageBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddBusiness(ctx, req.(*AddBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateBusiness(ctx, req.(*UpdateBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteBusiness(ctx, req.(*DeleteBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetBusinessValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBusinessValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetBusinessValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBusinessValues(ctx, req.(*GetBusinessValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateBusinessValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusinessValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateBusinessValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateBusinessValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateBusinessValues(ctx, req.(*UpdateBusinessValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configure_business.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PageBusiness",
			Handler:    _Service_PageBusiness_Handler,
		},
		{
			MethodName: "AddBusiness",
			Handler:    _Service_AddBusiness_Handler,
		},
		{
			MethodName: "UpdateBusiness",
			Handler:    _Service_UpdateBusiness_Handler,
		},
		{
			MethodName: "DeleteBusiness",
			Handler:    _Service_DeleteBusiness_Handler,
		},
		{
			MethodName: "GetBusinessValues",
			Handler:    _Service_GetBusinessValues_Handler,
		},
		{
			MethodName: "UpdateBusinessValues",
			Handler:    _Service_UpdateBusinessValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configure_business_service.proto",
}