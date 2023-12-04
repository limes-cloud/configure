package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllResourceValue(ctx context.Context, in *v1.AllResourceValueRequest) (*v1.AllResourceValueReply, error) {
	return s.ResourceValue.All(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateResourceValue(ctx context.Context, in *v1.UpdateResourceValueRequest) (*emptypb.Empty, error) {
	return s.ResourceValue.Update(kratosx.MustContext(ctx), in)
}
