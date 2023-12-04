package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
)

func (s *Service) AllResourceServer(ctx context.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	return s.ResourceServer.AllServer(kratosx.MustContext(ctx), in)
}

func (s *Service) PageServerResource(ctx context.Context, in *v1.PageServerResourceRequest) (*v1.PageServerResourceReply, error) {
	return s.ResourceServer.PageResource(kratosx.MustContext(ctx), in)
}
