package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/internal/core"
	kmodel "github.com/limes-cloud/kratosx/model"
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
	srv       *service.Template
	namespace *service.TemplateNamespace
}

func NewTemplate() *Template {
	return &Template{
		srv: service.NewTemplate(
			dbs.NewTemplate(),
			dbs.NewTemplateNamespace(),
			dbs.NewBusiness(),
			dbs.NewResource(),
		),
		namespace: service.NewTemplateNamespace(dbs.NewTemplateNamespace(), dbs.NewTemplate()),
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

func (s *Template) ListTemplateNamespace(ctx context.Context, req *pb.ListTemplateNamespaceRequest) (*pb.ListTemplateNamespaceResponse, error) {
	list, total, err := s.namespace.ListTemplateNamespace(core.MustContext(ctx), &types.ListTemplateNamespaceRequest{
		Page:      req.Page,
		PageSize:  req.PageSize,
		App:       req.App,
		Namespace: req.Namespace,
		Name:      req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListTemplateNamespaceResponse{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) CreateTemplateNamespace(ctx context.Context, req *pb.CreateTemplateNamespaceRequest) (*pb.CreateTemplateNamespaceResponse, error) {
	id, err := s.namespace.CreateTemplateNamespace(core.MustContext(ctx), &entity.TemplateNamespace{
		App:         req.App,
		Namespace:   req.Namespace,
		Name:        req.Name,
		Description: req.Description,
		Sort:        req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTemplateNamespaceResponse{Id: id}, nil
}

func (s *Template) UpdateTemplateNamespace(ctx context.Context, req *pb.UpdateTemplateNamespaceRequest) (*pb.UpdateTemplateNamespaceResponse, error) {
	if err := s.namespace.UpdateTemplateNamespace(core.MustContext(ctx), &entity.TemplateNamespace{
		BaseModel:   kmodel.BaseModel{Id: req.Id},
		Namespace:   req.Namespace,
		Name:        req.Name,
		Description: req.Description,
		Sort:        req.Sort,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateTemplateNamespaceResponse{}, nil
}

func (s *Template) DeleteTemplateNamespace(ctx context.Context, req *pb.DeleteTemplateNamespaceRequest) (*pb.DeleteTemplateNamespaceResponse, error) {
	if err := s.namespace.DeleteTemplateNamespace(core.MustContext(ctx), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteTemplateNamespaceResponse{}, nil
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
