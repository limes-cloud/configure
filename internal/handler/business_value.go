package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllBusinessValue(ctx context.Context, in *v1.AllBusinessValueRequest) (*v1.AllBusinessValueReply, error) {
	return s.BusinessValue.All(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateBusinessValue(ctx context.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	return s.BusinessValue.Update(kratosx.MustContext(ctx), in)
}
