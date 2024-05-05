package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/configure/api/business/v1"
	"github.com/limes-cloud/configure/api/errors"
	biz "github.com/limes-cloud/configure/internal/biz/business"
	"github.com/limes-cloud/configure/internal/config"
	data "github.com/limes-cloud/configure/internal/data/business"
)

type BusinessService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewBusinessService(conf *config.Config) *BusinessService {
	return &BusinessService{
		uc: biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *BusinessService) PageBusiness(ctx context.Context, in *v1.PageBusinessRequest) (*v1.PageBusinessReply, error) {
	req := &biz.PageBusinessRequest{}
	if err := util.Transform(in, &req); err != nil {
		return nil, errors.TransformError()
	}

	resources, total, err := s.uc.PageBusiness(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageBusinessReply{Total: total}
	if err := util.Transform(resources, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *BusinessService) AddBusiness(ctx context.Context, in *v1.AddBusinessRequest) (*emptypb.Empty, error) {
	resource := &biz.Business{}
	if err := util.Transform(in, resource); err != nil {
		return nil, errors.TransformError()
	}
	_, err := s.uc.AddBusiness(kratosx.MustContext(ctx), resource)
	return nil, err
}

func (s *BusinessService) UpdateBusiness(ctx context.Context, in *v1.UpdateBusinessRequest) (*emptypb.Empty, error) {
	resource := biz.Business{}
	if err := util.Transform(in, &resource); err != nil {
		return nil, errors.TransformError()
	}
	return nil, s.uc.UpdateBusiness(kratosx.MustContext(ctx), &resource)
}

func (s *BusinessService) DeleteBusiness(ctx context.Context, in *v1.DeleteBusinessRequest) (*emptypb.Empty, error) {
	return nil, s.uc.DeleteBusiness(kratosx.MustContext(ctx), in.Id)
}

func (s *BusinessService) GetBusinessValues(ctx context.Context, in *v1.GetBusinessValuesRequest) (*v1.GetBusinessValuesReply, error) {
	list, err := s.uc.GetBusinessValues(kratosx.MustContext(ctx), in.BusinessId)
	if err != nil {
		return nil, err
	}
	reply := v1.GetBusinessValuesReply{}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *BusinessService) UpdateBusinessValues(ctx context.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	for _, item := range in.List {
		rv := biz.BusinessValue{
			BusinessId: in.BusinessId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		}
		if err := s.uc.UpdateBusinessValue(kratosx.MustContext(ctx), &rv); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
