package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) PageResource(ctx context.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	return s.Resource.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddResource(ctx context.Context, in *v1.AddResourceRequest) (*emptypb.Empty, error) {
	return s.Resource.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateResource(ctx context.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	return s.Resource.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteResource(ctx context.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	return s.Resource.Delete(kratosx.MustContext(ctx), in)
}
