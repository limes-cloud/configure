package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetServer(ctx context.Context, in *v1.GetServerRequest) (*v1.GetServerReply, error) {
	return s.Server.Get(kratos.MustContext(ctx), in)
}

func (s *Service) PageServer(ctx context.Context, in *v1.PageServerRequest) (*v1.PageServerReply, error) {
	return s.Server.Page(kratos.MustContext(ctx), in)
}

func (s *Service) AddServer(ctx context.Context, in *v1.AddServerRequest) (*emptypb.Empty, error) {
	return s.Server.Add(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateServer(ctx context.Context, in *v1.UpdateServerRequest) (*emptypb.Empty, error) {
	return s.Server.Update(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteServer(ctx context.Context, in *v1.DeleteServerRequest) (*emptypb.Empty, error) {
	return s.Server.Delete(kratos.MustContext(ctx), in)
}
