package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) PageResource(ctx context.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	return s.logic.PageResource(kratos.MustContext(ctx), in)
}

func (s *Service) AddResource(ctx context.Context, in *v1.AddResourceRequest) (*emptypb.Empty, error) {
	return s.logic.AddResource(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateResource(ctx context.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	return s.logic.UpdateResource(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteResource(ctx context.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	return s.logic.DeleteResource(kratos.MustContext(ctx), in)
}
