package business

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetBusiness(ctx kratosx.Context, id uint32) (*Business, error)
	GetBusinessByKeyword(ctx kratosx.Context, key string) (*Business, error)
	PageBusiness(ctx kratosx.Context, req *PageBusinessRequest) ([]*Business, uint32, error)

	// AllBusiness(ctx kratosx.Context, id uint32) ([]*Business, error)
	AddBusiness(ctx kratosx.Context, c *Business) (uint32, error)

	UpdateBusiness(ctx kratosx.Context, c *Business) error
	DeleteBusiness(ctx kratosx.Context, uint322 uint32) error
	GetBusinessValues(ctx kratosx.Context, bid uint32) ([]*BusinessValue, error)
	UpdateBusinessValue(ctx kratosx.Context, rvs *BusinessValue) error
	CheckBusinessValue(ctx kratosx.Context, id uint32, value string) bool
	AllBusinessField(ctx kratosx.Context, sid uint32) ([]string, error)
	AllBusinessValue(ctx kratosx.Context, eid, sid uint32) ([]*BusinessValue, error)
}
