package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/configure/api/errors"
	v1 "github.com/limes-cloud/configure/api/server/v1"
	biz "github.com/limes-cloud/configure/internal/biz/server"
	"github.com/limes-cloud/configure/internal/config"
	data "github.com/limes-cloud/configure/internal/data/server"
)

type ServerService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewServerService(conf *config.Config) *ServerService {
	return &ServerService{
		uc: biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *ServerService) GetServer(ctx context.Context, in *v1.GetServerRequest) (*v1.Server, error) {
	server, err := s.uc.GetServer(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.Server{}
	if err := util.Transform(server, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ServerService) PageServer(ctx context.Context, in *v1.PageServerRequest) (*v1.PageServerReply, error) {
	req := &biz.PageServerRequest{}
	if err := util.Transform(in, &req); err != nil {
		return nil, errors.TransformError()
	}

	servers, total, err := s.uc.PageServer(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageServerReply{Total: total}
	if err := util.Transform(servers, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ServerService) AddServer(ctx context.Context, in *v1.AddServerRequest) (*v1.AddServerReply, error) {
	server := biz.Server{}
	if err := util.Transform(in, &server); err != nil {
		return nil, errors.TransformError()
	}
	id, err := s.uc.AddServer(kratosx.MustContext(ctx), &server)
	return &v1.AddServerReply{Id: id}, err
}

func (s *ServerService) UpdateServer(ctx context.Context, in *v1.UpdateServerRequest) (*emptypb.Empty, error) {
	server := &biz.Server{}
	if err := util.Transform(in, &server); err != nil {
		return nil, errors.TransformError()
	}
	return nil, s.uc.UpdateServer(kratosx.MustContext(ctx), server)
}

func (s *ServerService) DeleteServer(ctx context.Context, in *v1.DeleteServerRequest) (*emptypb.Empty, error) {
	return nil, s.uc.DeleteServer(kratosx.MustContext(ctx), in.Id)
}
