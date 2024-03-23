package server

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetServer(ctx kratosx.Context, id uint32) (*Server, error)
	GetServerByKeyword(ctx kratosx.Context, key string) (*Server, error)
	PageServer(ctx kratosx.Context, req *PageServerRequest) ([]*Server, uint32, error)
	AddServer(ctx kratosx.Context, c *Server) (uint32, error)
	UpdateServer(ctx kratosx.Context, c *Server) error
	DeleteServer(ctx kratosx.Context, uint322 uint32) error
}
