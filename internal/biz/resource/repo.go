package resource

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetResource(ctx kratosx.Context, id uint32) (*Resource, error)
	GetResourceByKeyword(ctx kratosx.Context, key string) (*Resource, error)
	PageResource(ctx kratosx.Context, req *PageResourceRequest) ([]*Resource, uint32, error)
	AddResource(ctx kratosx.Context, c *Resource) (uint32, error)
	UpdateResource(ctx kratosx.Context, c *Resource) error
	DeleteResource(ctx kratosx.Context, uint322 uint32) error
	GetResourceValues(ctx kratosx.Context, rid uint32) ([]*ResourceValue, error)
	UpdateResourceValue(ctx kratosx.Context, rv *ResourceValue) error
	GetResourceServerIds(ctx kratosx.Context, id uint32) ([]uint32, error)
	AllResourceValue(ctx kratosx.Context, eid, sid uint32) ([]*ResourceValue, error)
	AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error)
	CheckResourceValue(ctx kratosx.Context, rid uint32, value string) error
}
