package service

import (
	"context"

	"github.com/limes-cloud/configure/internal/biz/types"

	"github.com/limes-cloud/configure/internal/biz"

	"github.com/jinzhu/copier"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CurrentTemplate(ctx context.Context, in *v1.CurrentTemplateRequest) (*v1.CurrentTemplateReply, error) {
	template, err := s.TemplateUseCase.Current(kratosx.MustContext(ctx), in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := v1.CurrentTemplateReply{}
	if err := copier.Copy(&reply, template); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) GetTemplate(ctx context.Context, in *v1.GetTemplateRequest) (*v1.GetTemplateReply, error) {
	template, err := s.TemplateUseCase.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.GetTemplateReply{}
	if err := copier.Copy(&reply, template); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) PageTemplate(ctx context.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	req := &types.PageTemplateRequest{}
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	templates, total, err := s.TemplateUseCase.Page(kratosx.MustContext(ctx), req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageTemplateReply{Total: total}
	if err := copier.Copy(&reply.List, templates); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddTemplate(ctx context.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	template := &biz.Template{}
	if err := copier.Copy(template, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.TemplateUseCase.Add(kratosx.MustContext(ctx), template)
	return nil, err
}

func (s *Service) SwitchTemplate(ctx context.Context, in *v1.SwitchTemplateRequest) (*emptypb.Empty, error) {
	return nil, s.TemplateUseCase.Switch(kratosx.MustContext(ctx), in.ServerId, in.Id)
}

func (s *Service) CompareTemplate(ctx context.Context, in *v1.CompareTemplateRequest) (*v1.CompareTemplateReply, error) {
	list, err := s.TemplateUseCase.Compare(kratosx.MustContext(ctx), &types.CompareTemplateRequest{
		Id:      in.Id,
		Format:  in.Format,
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}
	reply := v1.CompareTemplateReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) ParseTemplatePreview(ctx context.Context, in *v1.ParseTemplatePreviewRequest) (*v1.ParseTemplatePreviewReply, error) {
	content, err := s.TemplateUseCase.ParseByContent(kratosx.MustContext(ctx), &types.ParseByContentRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
		Format:   in.Format,
		Content:  in.Content,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ParseTemplatePreviewReply{Content: content}, nil
}

func (s *Service) ParseTemplate(ctx context.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	res, err := s.TemplateUseCase.Parse(kratosx.MustContext(ctx), &types.ParseRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ParseTemplateReply{Content: res.Content, Format: res.Format}, nil
}
