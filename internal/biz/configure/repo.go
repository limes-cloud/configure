package configure

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetConfigure(ctx kratosx.Context, id uint32) (*Configure, error)
	GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*Configure, error)
	ListConfigure(ctx kratosx.Context, req *ListConfigureRequest) ([]*Configure, uint32, error)
	CreateConfigure(ctx kratosx.Context, c *Configure) (uint32, error)
	UpdateConfigure(ctx kratosx.Context, c *Configure) error
	CurrentTemplate(ctx kratosx.Context, srvId uint32) (string, string, error)
	GetServerIdByKeyword(ctx kratosx.Context, keyword string) (uint32, error)
	GetEnvIdByToken(ctx kratosx.Context, token string) (uint32, error)
}
