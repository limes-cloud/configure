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

func (s *Service) PageResource(ctx context.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	req := &types.PageResourceRequest{}
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	resources, total, err := s.ResourceUseCase.Page(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageResourceReply{Total: total}
	if err := copier.Copy(&reply.List, resources); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddResource(ctx context.Context, in *v1.AddResourceRequest) (*emptypb.Empty, error) {
	resource := &biz.Resource{}
	if err := copier.Copy(resource, in); err != nil {
		return nil, v1.TransformError()
	}
	if in.Private != nil && *in.Private == true {
		for _, srvId := range in.Servers {
			resource.ResourceServer = append(resource.ResourceServer, &biz.ResourceServer{
				ServerID: srvId,
			})
		}
	}
	_, err := s.ResourceUseCase.Add(kratosx.MustContext(ctx), resource)
	return nil, err
}

func (s *Service) UpdateResource(ctx context.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	resource := &biz.Resource{}
	if err := copier.Copy(resource, in); err != nil {
		return nil, v1.TransformError()
	}
	if in.Private != nil && *in.Private == true {
		for _, srvId := range in.Servers {
			resource.ResourceServer = append(resource.ResourceServer, &biz.ResourceServer{
				ServerID: srvId,
			})
		}
	}
	return nil, s.ResourceUseCase.Update(kratosx.MustContext(ctx), resource)
}

func (s *Service) DeleteResource(ctx context.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	return nil, s.ResourceUseCase.Delete(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) AllResourceServer(ctx context.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	kCtx := kratosx.MustContext(ctx)
	ids, err := s.ResourceUseCase.AllResourceServerIds(kCtx, in.ResourceId)
	if err != nil {
		return nil, err
	}
	envs, err := s.ServerUseCase.GetByIds(kCtx, ids)
	if err != nil {
		return nil, err
	}
	reply := v1.AllResourceServerReply{}
	if err := copier.Copy(&reply.List, envs); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) PageServerResource(ctx context.Context, in *v1.PageServerResourceRequest) (*v1.PageServerResourceReply, error) {
	list, total, err := s.ResourceUseCase.PageServerResource(
		kratosx.MustContext(ctx),
		&types.PageServerResourceRequest{
			Page:     in.Page,
			PageSize: in.PageSize,
			ServerID: in.ServerId,
		})
	if err != nil {
		return nil, err
	}

	reply := v1.PageServerResourceReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) AllResourceValue(ctx context.Context, in *v1.AllResourceValueRequest) (*v1.AllResourceValueReply, error) {
	list, err := s.ResourceUseCase.AllResourceValue(kratosx.MustContext(ctx), in.ResourceId)
	if err != nil {
		return nil, err
	}
	reply := v1.AllResourceValueReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) UpdateResourceValue(ctx context.Context, in *v1.UpdateResourceValueRequest) (*emptypb.Empty, error) {
	for _, item := range in.List {
		rv := biz.ResourceValue{
			ResourceID: in.ResourceId,
			EnvID:      item.EnvId,
			Value:      item.Value,
		}
		if err := s.ResourceUseCase.UpdateResourceValue(kratosx.MustContext(ctx), &rv); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
