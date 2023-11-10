package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"
)

func (s *Service) AllResourceServer(ctx context.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	return s.ResourceServer.AllServer(kratos.MustContext(ctx), in)
}

func (s *Service) PageServerResource(ctx context.Context, in *v1.PageServerResourceRequest) (*v1.PageServerResourceReply, error) {
	return s.ResourceServer.PageResource(kratos.MustContext(ctx), in)
}
