package service

import (
	"context"

	"github.com/limes-cloud/configure/internal/biz/types"

	"github.com/limes-cloud/configure/internal/biz"

	"github.com/jinzhu/copier"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetServer(ctx context.Context, in *v1.GetServerRequest) (*v1.GetServerReply, error) {
	server, err := s.ServerUseCase.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.GetServerReply{}
	if err := copier.Copy(&reply, server); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) PageServer(ctx context.Context, in *v1.PageServerRequest) (*v1.PageServerReply, error) {
	req := &types.PageServerRequest{}
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	servers, total, err := s.ServerUseCase.Page(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageServerReply{Total: total}
	if err := copier.Copy(&reply.List, servers); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddServer(ctx context.Context, in *v1.AddServerRequest) (*emptypb.Empty, error) {
	server := &biz.Server{}
	if err := copier.Copy(server, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.ServerUseCase.Add(kratosx.MustContext(ctx), server)
	return nil, err
}

func (s *Service) UpdateServer(ctx context.Context, in *v1.UpdateServerRequest) (*emptypb.Empty, error) {
	server := &biz.Server{}
	if err := copier.Copy(server, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.ServerUseCase.Update(kratosx.MustContext(ctx), server)
}

func (s *Service) DeleteServer(ctx context.Context, in *v1.DeleteServerRequest) (*emptypb.Empty, error) {
	return nil, s.ServerUseCase.Delete(kratosx.MustContext(ctx), in.Id)
}
