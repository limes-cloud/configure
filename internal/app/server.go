package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	kmodel "github.com/limes-cloud/kratosx/model"

	pb "github.com/limes-cloud/configure/api/server"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
	"github.com/limes-cloud/configure/internal/infra/rpc"
	"github.com/limes-cloud/configure/internal/types"
)

type Server struct {
	pb.UnimplementedServerServer
	srv *service.Server
}

func NewServer() *Server {
	return &Server{
		srv: service.NewServer(
			dbs.NewServer(),
			rpc.NewPermission(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewServer()
		pb.RegisterServerHTTPServer(hs, srv)
		pb.RegisterServerServer(gs, srv)
	})
}

// ListServer 获取服务信息列表
func (s *Server) ListServer(c context.Context, req *pb.ListServerRequest) (*pb.ListServerResponse, error) {
	list, total, err := s.srv.ListServer(kratosx.MustContext(c), &types.ListServerRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListServerResponse{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListServerResponse_Server{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Name:        item.Name,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateServer 创建服务信息
func (s *Server) CreateServer(c context.Context, req *pb.CreateServerRequest) (*pb.CreateServerResponse, error) {
	id, err := s.srv.CreateServer(kratosx.MustContext(c), &entity.Server{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateServerResponse{Id: id}, nil
}

// UpdateServer 更新服务信息
func (s *Server) UpdateServer(c context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerResponse, error) {
	if err := s.srv.UpdateServer(kratosx.MustContext(c), &entity.Server{
		BaseModel:   kmodel.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateServerResponse{}, nil
}

// DeleteServer 删除服务信息
func (s *Server) DeleteServer(c context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerResponse, error) {
	if err := s.srv.DeleteServer(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteServerResponse{}, nil
}
