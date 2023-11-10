package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllResourceValue(ctx context.Context, in *v1.AllResourceValueRequest) (*v1.AllResourceValueReply, error) {
	return s.ResourceValue.All(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateResourceValue(ctx context.Context, in *v1.UpdateResourceValueRequest) (*emptypb.Empty, error) {
	return s.ResourceValue.Update(kratos.MustContext(ctx), in)
}
