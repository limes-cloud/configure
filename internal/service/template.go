package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/configure/api/errors"
	v1 "github.com/limes-cloud/configure/api/template/v1"
	biz "github.com/limes-cloud/configure/internal/biz/template"
	"github.com/limes-cloud/configure/internal/config"
	"github.com/limes-cloud/configure/internal/data/business"
	"github.com/limes-cloud/configure/internal/data/resource"
	data "github.com/limes-cloud/configure/internal/data/template"
)

type TemplateService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewTemplateService(conf *config.Config) *TemplateService {
	return &TemplateService{
		uc: biz.NewUseCase(conf, data.NewRepo(), resource.NewRepo(), business.NewRepo()),
	}
}

func (s *TemplateService) CurrentTemplate(ctx context.Context, in *v1.CurrentTemplateRequest) (*v1.CurrentTemplateReply, error) {
	template, err := s.uc.CurrentTemplate(kratosx.MustContext(ctx), in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := v1.CurrentTemplateReply{}
	if err := copier.Copy(&reply, template); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) GetTemplate(ctx context.Context, in *v1.GetTemplateRequest) (*v1.GetTemplateReply, error) {
	template, err := s.uc.GetTemplate(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.GetTemplateReply{}
	if err := copier.Copy(&reply, template); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) PageTemplate(ctx context.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	req := &biz.PageTemplateRequest{}
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.TransformError()
	}

	templates, total, err := s.uc.PageTemplate(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageTemplateReply{Total: total}
	if err := copier.Copy(&reply.List, templates); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) AddTemplate(ctx context.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	template := &biz.Template{}
	if err := copier.Copy(template, in); err != nil {
		return nil, errors.TransformError()
	}
	_, err := s.uc.AddTemplate(kratosx.MustContext(ctx), template)
	return nil, err
}

func (s *TemplateService) SwitchTemplate(ctx context.Context, in *v1.SwitchTemplateRequest) (*emptypb.Empty, error) {
	return nil, s.uc.SwitchTemplate(kratosx.MustContext(ctx), in.ServerId, in.Id)
}

func (s *TemplateService) CompareTemplate(ctx context.Context, in *v1.CompareTemplateRequest) (*v1.CompareTemplateReply, error) {
	list, err := s.uc.Compare(kratosx.MustContext(ctx), &biz.CompareTemplateRequest{
		Id:      in.Id,
		Format:  in.Format,
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}
	reply := v1.CompareTemplateReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) PreviewTemplate(ctx context.Context, in *v1.PreviewTemplateRequest) (*v1.PreviewTemplateReply, error) {
	res, err := s.uc.PreviewTemplate(kratosx.MustContext(ctx), &biz.PreviewTemplateRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
		Format:   in.Format,
		Content:  in.Content,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PreviewTemplateReply{Content: res.Content}, nil
}

func (s *TemplateService) ParseTemplate(ctx context.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	res, err := s.uc.ParseTemplate(kratosx.MustContext(ctx), &biz.ParseTemplateRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ParseTemplateReply{Content: res.Content, Format: res.Format}, nil
}
