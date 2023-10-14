package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) PageBusiness(ctx context.Context, in *v1.PageBusinessRequest) (*v1.PageBusinessReply, error) {
	return s.Business.Page(kratos.MustContext(ctx), in)
}

func (s *Service) AddBusiness(ctx context.Context, in *v1.AddBusinessRequest) (*emptypb.Empty, error) {
	return s.Business.Add(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateBusiness(ctx context.Context, in *v1.UpdateBusinessRequest) (*emptypb.Empty, error) {
	return s.Business.Update(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteBusiness(ctx context.Context, in *v1.DeleteBusinessRequest) (*emptypb.Empty, error) {
	return s.Business.Delete(kratos.MustContext(ctx), in)
}
