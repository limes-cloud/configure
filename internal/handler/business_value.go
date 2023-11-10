package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllBusinessValue(ctx context.Context, in *v1.AllBusinessValueRequest) (*v1.AllBusinessValueReply, error) {
	return s.BusinessValue.All(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateBusinessValue(ctx context.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	return s.BusinessValue.Update(kratos.MustContext(ctx), in)
}
