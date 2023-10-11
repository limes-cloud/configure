package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllResourceServer(ctx context.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	return s.logic.AllResourceServer(kratos.MustContext(ctx), in)
}

func (s *Service) AddResourceServer(ctx context.Context, in *v1.AddResourceServerRequest) (*emptypb.Empty, error) {
	return s.logic.AddResourceServer(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteResourceServer(ctx context.Context, in *v1.DeleteResourceServerRequest) (*emptypb.Empty, error) {
	return s.logic.DeleteResourceServer(kratos.MustContext(ctx), in)
}
