package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/configure/api/configure/errors"
	pb "github.com/limes-cloud/configure/api/configure/template/v1"
	"github.com/limes-cloud/configure/internal/biz/template"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/data"
	"github.com/limes-cloud/configure/internal/factory"
)

type TemplateService struct {
	pb.UnimplementedTemplateServer
	uc *template.UseCase
}

func NewTemplateService(conf *conf.Config) *TemplateService {
	return &TemplateService{
		uc: template.NewUseCase(
			conf,
			data.NewTemplateRepo(),
			factory.New(data.NewBusinessRepo(), data.NewResourceRepo()),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewTemplateService(c)
		pb.RegisterTemplateHTTPServer(hs, srv)
		pb.RegisterTemplateServer(gs, srv)
	})
}

func (s *TemplateService) CurrentTemplate(ctx context.Context, req *pb.CurrentTemplateRequest) (*pb.CurrentTemplateReply, error) {
	res, err := s.uc.CurrentTemplate(kratosx.MustContext(ctx), req.ServerId)
	if err != nil {
		return nil, err
	}
	reply := pb.CurrentTemplateReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateReply, error) {
	res, err := s.uc.GetTemplate(kratosx.MustContext(ctx), req.Id)
	if err != nil {
		return nil, err
	}
	reply := pb.GetTemplateReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	res, total, err := s.uc.ListTemplate(kratosx.MustContext(ctx), &template.ListTemplateRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListTemplateReply{Total: total}
	if err := valx.Transform(res, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) CreateTemplate(c context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateReply, error) {
	var (
		in  = template.Template{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	id, err := s.uc.CreateTemplate(kratosx.MustContext(ctx), &in)
	return &pb.CreateTemplateReply{Id: id}, err
}

func (s *TemplateService) SwitchTemplate(ctx context.Context, req *pb.SwitchTemplateRequest) (*pb.SwitchTemplateReply, error) {
	return &pb.SwitchTemplateReply{}, s.uc.SwitchTemplate(kratosx.MustContext(ctx), req.ServerId, req.Id)
}

func (s *TemplateService) CompareTemplate(ctx context.Context, req *pb.CompareTemplateRequest) (*pb.CompareTemplateReply, error) {
	list, err := s.uc.CompareTemplate(kratosx.MustContext(ctx), &template.CompareTemplateRequest{
		Id:      req.Id,
		Format:  req.Format,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CompareTemplateReply{}
	if err := valx.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *TemplateService) PreviewTemplate(ctx context.Context, req *pb.PreviewTemplateRequest) (*pb.PreviewTemplateReply, error) {
	res, err := s.uc.PreviewTemplate(kratosx.MustContext(ctx), &template.PreviewTemplateRequest{
		EnvId:    req.EnvId,
		ServerId: req.ServerId,
		Format:   req.Format,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PreviewTemplateReply{Content: res.Content}, nil
}

func (s *TemplateService) ParseTemplate(ctx context.Context, req *pb.ParseTemplateRequest) (*pb.ParseTemplateReply, error) {
	res, err := s.uc.ParseTemplate(kratosx.MustContext(ctx), &template.ParseTemplateRequest{
		EnvId:    req.EnvId,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ParseTemplateReply{Content: res.Content, Format: res.Format}, nil
}
