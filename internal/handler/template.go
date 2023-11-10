package handler

import (
	"context"

	"github.com/limes-cloud/kratos"

	v1 "github.com/limes-cloud/configure/api/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CurrentTemplate(ctx context.Context, in *v1.CurrentTemplateRequest) (*v1.CurrentTemplateReply, error) {
	return s.Template.Current(kratos.MustContext(ctx), in)
}

func (s *Service) GetTemplate(ctx context.Context, in *v1.GetTemplateRequest) (*v1.GetTemplateReply, error) {
	return s.Template.Get(kratos.MustContext(ctx), in)
}

func (s *Service) PageTemplate(ctx context.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	return s.Template.Page(kratos.MustContext(ctx), in)
}

func (s *Service) AddTemplate(ctx context.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	return s.Template.Add(kratos.MustContext(ctx), in)
}

func (s *Service) UseTemplateVersion(ctx context.Context, in *v1.UseTemplateVersionRequest) (*emptypb.Empty, error) {
	return s.Template.UseVersion(kratos.MustContext(ctx), in)
}

func (s *Service) ParseTemplatePreview(ctx context.Context, in *v1.ParseTemplatePreviewRequest) (*v1.ParseTemplatePreviewReply, error) {
	return s.Template.ParsePreview(kratos.MustContext(ctx), in)
}

func (s *Service) ParseTemplate(ctx context.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	return s.Template.Parse(kratos.MustContext(ctx), in)
}
