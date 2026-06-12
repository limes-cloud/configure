package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/configure/api/errors"
	pb "github.com/limes-cloud/configure/api/template"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
	"github.com/limes-cloud/configure/internal/types"
)

type Template struct {
	pb.UnimplementedTemplateServer
	srv *service.Template
}

func NewTemplate() *Template {
	return &Template{
		srv: service.NewTemplate(
			dbs.NewTemplate(),
			dbs.NewBusiness(),
			dbs.NewResource(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewTemplate()
		pb.RegisterTemplateHTTPServer(hs, srv)
		pb.RegisterTemplateServer(gs, srv)
	})
}

func (s *Template) CurrentTemplate(ctx context.Context, req *pb.CurrentTemplateRequest) (*pb.CurrentTemplateResponse, error) {
	res, err := s.srv.CurrentTemplate(core.MustContext(ctx), req.App, req.GetNamespace())
	if err != nil {
		return nil, err
	}
	reply := pb.CurrentTemplateResponse{}
	if err := value.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
	res, err := s.srv.GetTemplate(core.MustContext(ctx), req.Id)
	if err != nil {
		return nil, err
	}
	reply := pb.GetTemplateResponse{}
	if err := value.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateResponse, error) {
	res, total, err := s.srv.ListTemplate(core.MustContext(ctx), &types.ListTemplateRequest{
		Page:      req.Page,
		PageSize:  req.PageSize,
		App:       req.App,
		Namespace: req.Namespace,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListTemplateResponse{Total: total}
	if err := value.Transform(res, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) CreateTemplate(c context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateResponse, error) {
	id, err := s.srv.CreateTemplate(core.MustContext(c), &entity.Template{
		App:         req.App,
		Namespace:   req.GetNamespace(),
		Content:     req.Content,
		Description: req.Description,
		Format:      req.Format,
	})
	return &pb.CreateTemplateResponse{Id: id}, err
}

func (s *Template) SwitchTemplate(ctx context.Context, req *pb.SwitchTemplateRequest) (*pb.SwitchTemplateResponse, error) {
	return &pb.SwitchTemplateResponse{}, s.srv.SwitchTemplate(core.MustContext(ctx), req.App, req.GetNamespace(), req.Id)
}

func (s *Template) CompareTemplate(ctx context.Context, req *pb.CompareTemplateRequest) (*pb.CompareTemplateResponse, error) {
	list, err := s.srv.CompareTemplate(core.MustContext(ctx), &types.CompareTemplateRequest{
		Id:      req.Id,
		Format:  req.Format,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CompareTemplateResponse{}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) PreviewTemplate(ctx context.Context, req *pb.PreviewTemplateRequest) (*pb.PreviewTemplateResponse, error) {
	res, err := s.srv.PreviewTemplate(core.MustContext(ctx), &types.PreviewTemplateRequest{
		EnvId:     req.EnvId,
		App:       req.App,
		Namespace: req.GetNamespace(),
		Format:    req.Format,
		Content:   req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PreviewTemplateResponse{Content: res.Content}, nil
}
