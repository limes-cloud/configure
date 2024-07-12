package server

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetServer 获取指定的服务信息
	GetServer(ctx kratosx.Context, id uint32) (*Server, error)

	// ListServer 获取服务信息列表
	ListServer(ctx kratosx.Context, req *ListServerRequest) ([]*Server, uint32, error)

	// CreateServer 创建服务信息
	CreateServer(ctx kratosx.Context, req *Server) (uint32, error)

	// UpdateServer 更新服务信息
	UpdateServer(ctx kratosx.Context, req *Server) error

	// DeleteServer 删除服务信息
	DeleteServer(ctx kratosx.Context, id uint32) error

	// GetServerByKeyword 获取指定的服务信息
	GetServerByKeyword(ctx kratosx.Context, keyword string) (*Server, error)

	// UpdateServerStatus 更新服务信息状态
	UpdateServerStatus(ctx kratosx.Context, id uint32, status bool) error
}
