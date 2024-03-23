package configure

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Repo interface {
	GetConfigure(ctx kratosx.Context, id uint32) (*Configure, error)
	GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*Configure, error)
	PageConfigure(ctx kratosx.Context, options *ktypes.PageOptions) ([]*Configure, uint32, error)
	AddConfigure(ctx kratosx.Context, c *Configure) (uint32, error)
	UpdateConfigure(ctx kratosx.Context, c *Configure) error
}
