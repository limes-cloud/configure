package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/configure/api/errors"
	v1 "github.com/limes-cloud/configure/api/resource/v1"
	biz "github.com/limes-cloud/configure/internal/biz/resource"
	"github.com/limes-cloud/configure/internal/config"
	data "github.com/limes-cloud/configure/internal/data/resource"
)

type ResourceService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewResourceService(conf *config.Config) *ResourceService {
	return &ResourceService{
		uc: biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *ResourceService) PageResource(ctx context.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	req := &biz.PageResourceRequest{}
	if err := util.Transform(in, &req); err != nil {
		return nil, errors.TransformError()
	}

	res, total, err := s.uc.PageResource(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageResourceReply{Total: total}
	if err := util.Transform(res, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	for ind := range reply.List {
		var ids []uint32
		for _, ite := range res[ind].ResourceServers {
			ids = append(ids, ite.ServerId)
		}
		reply.List[ind].Servers = ids
	}

	return &reply, nil
}

func (s *ResourceService) AddResource(ctx context.Context, in *v1.AddResourceRequest) (*v1.AddResourceReply, error) {
	resource := biz.Resource{}
	if err := util.Transform(in, resource); err != nil {
		return nil, errors.TransformError()
	}
	if in.Private != nil && *in.Private {
		for _, srvId := range in.Servers {
			resource.ResourceServers = append(resource.ResourceServers, &biz.ResourceServer{
				ServerId: srvId,
			})
		}
	}
	id, err := s.uc.AddResource(kratosx.MustContext(ctx), &resource)
	return &v1.AddResourceReply{
		Id: id,
	}, err
}

func (s *ResourceService) UpdateResource(ctx context.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	resource := &biz.Resource{}
	if err := util.Transform(in, resource); err != nil {
		return nil, errors.TransformError()
	}
	if in.Private != nil && *in.Private {
		for _, srvId := range in.Servers {
			resource.ResourceServers = append(resource.ResourceServers, &biz.ResourceServer{
				ServerId: srvId,
			})
		}
	}
	return nil, s.uc.UpdateResource(kratosx.MustContext(ctx), resource)
}

func (s *ResourceService) DeleteResource(ctx context.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	return nil, s.uc.DeleteResource(kratosx.MustContext(ctx), in.Id)
}

func (s *ResourceService) GetResourceServerIds(ctx context.Context, in *v1.GetResourceServerIdsRequest) (*v1.GetResourceServerIdsReply, error) {
	kCtx := kratosx.MustContext(ctx)
	ids, err := s.uc.GetResourceServerIds(kCtx, in.ResourceId)
	if err != nil {
		return nil, err
	}

	return &v1.GetResourceServerIdsReply{List: ids}, nil
}

func (s *ResourceService) GetResourceValues(ctx context.Context, in *v1.GetResourceValuesRequest) (*v1.GetResourceValuesReply, error) {
	list, err := s.uc.GetResourceValues(kratosx.MustContext(ctx), in.ResourceId)
	if err != nil {
		return nil, err
	}
	reply := v1.GetResourceValuesReply{}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ResourceService) UpdateResourceValues(ctx context.Context, in *v1.UpdateResourceValuesRequest) (*emptypb.Empty, error) {
	for _, item := range in.List {
		rv := biz.ResourceValue{
			ResourceId: in.ResourceId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		}
		if err := s.uc.UpdateResourceValue(kratosx.MustContext(ctx), &rv); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
