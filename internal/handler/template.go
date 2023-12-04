package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"

	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Service) UseTemplateVersion(ctx context.Context, in *v1.UseTemplateVersionRequest) (*emptypb.Empty, error) {
	return s.Template.UseVersion(kratosx.MustContext(ctx), in)
}

func (s *Service) ParseTemplatePreview(ctx context.Context, in *v1.ParseTemplatePreviewRequest) (*v1.ParseTemplatePreviewReply, error) {
	return s.Template.ParsePreview(kratosx.MustContext(ctx), in)
}

func (s *Service) ParseTemplate(ctx context.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	return s.Template.Parse(kratosx.MustContext(ctx), in)
}
