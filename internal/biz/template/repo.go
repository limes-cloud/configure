package template

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetTemplate(ctx kratosx.Context, id uint32) (*Template, error)
	GetTemplateByVersion(ctx kratosx.Context, version string) (*Template, error)
	CurrentTemplate(ctx kratosx.Context, srvId uint32) (*Template, error)
	PageTemplate(ctx kratosx.Context, options *PageTemplateRequest) ([]*Template, uint32, error)
	AddTemplate(ctx kratosx.Context, c *Template) (uint32, error)
	UpdateTemplate(ctx kratosx.Context, c *Template) error
	UseTemplate(ctx kratosx.Context, srvId, tpId uint32) error
	DeleteTemplate(ctx kratosx.Context, uint322 uint32) error
}
