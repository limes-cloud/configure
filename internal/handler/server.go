package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetServer(ctx context.Context, in *v1.GetServerRequest) (*v1.GetServerReply, error) {
	return s.logic.GetServer(kratos.MustContext(ctx), in)
}

func (s *Service) PageServer(ctx context.Context, in *v1.PageServerRequest) (*v1.PageServerReply, error) {
	return s.logic.PageServer(kratos.MustContext(ctx), in)
}

func (s *Service) AddServer(ctx context.Context, in *v1.AddServerRequest) (*emptypb.Empty, error) {
	return s.logic.AddServer(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateServer(ctx context.Context, in *v1.UpdateServerRequest) (*emptypb.Empty, error) {
	return s.logic.UpdateServer(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteServer(ctx context.Context, in *v1.DeleteServerRequest) (*emptypb.Empty, error) {
	return s.logic.DeleteServer(kratos.MustContext(ctx), in)
}
