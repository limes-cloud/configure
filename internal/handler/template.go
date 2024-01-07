package handler

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/configure/api/v1"
)

func (s *Service) CurrentTemplate(ctx context.Context, in *v1.CurrentTemplateRequest) (*v1.CurrentTemplateReply, error) {
	return s.Template.Current(kratosx.MustContext(ctx), in)
}

func (s *Service) GetTemplate(ctx context.Context, in *v1.GetTemplateRequest) (*v1.GetTemplateReply, error) {
	return s.Template.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) PageTemplate(ctx context.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	return s.Template.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddTemplate(ctx context.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	return s.Template.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) SwitchTemplate(ctx context.Context, in *v1.SwitchTemplateRequest) (*emptypb.Empty, error) {
	return s.Template.Switch(kratosx.MustContext(ctx), in)
}

func (s *Service) CompareTemplate(ctx context.Context, in *v1.CompareTemplateRequest) (*v1.CompareTemplateReply, error) {
	return s.Template.Compare(kratosx.MustContext(ctx), in)
}

func (s *Service) ParseTemplatePreview(ctx context.Context, in *v1.ParseTemplatePreviewRequest) (*v1.ParseTemplatePreviewReply, error) {
	return s.Template.ParsePreview(kratosx.MustContext(ctx), in)
}

func (s *Service) ParseTemplate(ctx context.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	return s.Template.Parse(kratosx.MustContext(ctx), in)
}
