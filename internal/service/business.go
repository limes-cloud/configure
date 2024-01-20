package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/internal/biz/types"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) PageBusiness(ctx context.Context, in *v1.PageBusinessRequest) (*v1.PageBusinessReply, error) {
	req := &types.PageServerBusinessRequest{}
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	resources, total, err := s.BusinessUseCase.PageServerBusiness(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageBusinessReply{Total: total}
	if err := copier.Copy(&reply.List, resources); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddBusiness(ctx context.Context, in *v1.AddBusinessRequest) (*emptypb.Empty, error) {
	resource := &biz.Business{}
	if err := copier.Copy(resource, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.BusinessUseCase.Add(kratosx.MustContext(ctx), resource)
	return nil, err
}

func (s *Service) UpdateBusiness(ctx context.Context, in *v1.UpdateBusinessRequest) (*emptypb.Empty, error) {
	resource := &biz.Business{}
	if err := copier.Copy(resource, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.BusinessUseCase.Update(kratosx.MustContext(ctx), resource)
}

func (s *Service) DeleteBusiness(ctx context.Context, in *v1.DeleteBusinessRequest) (*emptypb.Empty, error) {
	return nil, s.BusinessUseCase.Delete(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) AllBusinessValue(ctx context.Context, in *v1.AllBusinessValueRequest) (*v1.AllBusinessValueReply, error) {
	list, err := s.BusinessUseCase.AllBusinessValue(kratosx.MustContext(ctx), in.BusinessId)
	if err != nil {
		return nil, err
	}
	reply := v1.AllBusinessValueReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) UpdateBusinessValue(ctx context.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	for _, item := range in.List {
		rv := biz.BusinessValue{
			BusinessID: in.BusinessId,
			EnvID:      item.EnvId,
			Value:      item.Value,
		}
		if err := s.BusinessUseCase.UpdateBusinessValue(kratosx.MustContext(ctx), &rv); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
