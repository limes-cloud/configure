package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllResourceServer(ctx context.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	return s.ResourceServer.AllServer(kratos.MustContext(ctx), in)
}

func (s *Service) AllServerResource(ctx context.Context, in *v1.AllServerResourceRequest) (*v1.AllServerResourceReply, error) {
	return s.ResourceServer.AllResource(kratos.MustContext(ctx), in)
}

func (s *Service) AddResourceServer(ctx context.Context, in *v1.AddResourceServerRequest) (*emptypb.Empty, error) {
	return s.ResourceServer.Add(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteResourceServer(ctx context.Context, in *v1.DeleteResourceServerRequest) (*emptypb.Empty, error) {
	return s.ResourceServer.Delete(kratos.MustContext(ctx), in)
}
