// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.4
// source: configure_configure_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationServiceCompareConfigure = "/configure_configure.Service/CompareConfigure"
const OperationServiceGetConfigure = "/configure_configure.Service/GetConfigure"
const OperationServiceUpdateConfigure = "/configure_configure.Service/UpdateConfigure"

type ServiceHTTPServer interface {
	CompareConfigure(context.Context, *CompareConfigureRequest) (*CompareConfigureReply, error)
	GetConfigure(context.Context, *GetConfigureRequest) (*GetConfigureReply, error)
	UpdateConfigure(context.Context, *UpdateConfigureRequest) (*emptypb.Empty, error)
}

func RegisterServiceHTTPServer(s *http.Server, srv ServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/configure/v1/configure", _Service_GetConfigure0_HTTP_Handler(srv))
	r.PUT("/configure/v1/configure", _Service_UpdateConfigure0_HTTP_Handler(srv))
	r.POST("/configure/v1/configure/compare", _Service_CompareConfigure0_HTTP_Handler(srv))
}

func _Service_GetConfigure0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetConfigureRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetConfigure)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetConfigure(ctx, req.(*GetConfigureRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetConfigureReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateConfigure0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateConfigureRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateConfigure)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateConfigure(ctx, req.(*UpdateConfigureRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_CompareConfigure0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompareConfigureRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceCompareConfigure)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CompareConfigure(ctx, req.(*CompareConfigureRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompareConfigureReply)
		return ctx.Result(200, reply)
	}
}

type ServiceHTTPClient interface {
	CompareConfigure(ctx context.Context, req *CompareConfigureRequest, opts ...http.CallOption) (rsp *CompareConfigureReply, err error)
	GetConfigure(ctx context.Context, req *GetConfigureRequest, opts ...http.CallOption) (rsp *GetConfigureReply, err error)
	UpdateConfigure(ctx context.Context, req *UpdateConfigureRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type ServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceHTTPClient(client *http.Client) ServiceHTTPClient {
	return &ServiceHTTPClientImpl{client}
}

func (c *ServiceHTTPClientImpl) CompareConfigure(ctx context.Context, in *CompareConfigureRequest, opts ...http.CallOption) (*CompareConfigureReply, error) {
	var out CompareConfigureReply
	pattern := "/configure/v1/configure/compare"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceCompareConfigure))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetConfigure(ctx context.Context, in *GetConfigureRequest, opts ...http.CallOption) (*GetConfigureReply, error) {
	var out GetConfigureReply
	pattern := "/configure/v1/configure"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetConfigure))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateConfigure(ctx context.Context, in *UpdateConfigureRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/configure/v1/configure"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateConfigure))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
