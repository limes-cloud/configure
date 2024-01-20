package biz

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/biz/types"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Business struct {
	ktypes.BaseModel
	ServerID    uint32           `json:"server_id" gorm:"not null;comment:服务id"`
	Keyword     string           `json:"keyword" gorm:"not null;type:char(32) binary;comment:变量标识"`
	Type        string           `json:"type" gorm:"not null;size:32;comment:变量类型"`
	Description string           `json:"description" gorm:"not null;size:128;comment:变量描述"`
	Server      *Server          `json:"server" gorm:"constraint:onDelete:cascade"`
	Values      []*BusinessValue `json:"business_value" gorm:"constraint:onDelete:cascade"`
}

type BusinessValue struct {
	ktypes.BaseModel
	EnvID      uint32    `json:"env_id" gorm:"not null;comment:环境id"`
	BusinessID uint32    `json:"business_id" gorm:"not null;comment:业务变量id"`
	Value      string    `json:"value" gorm:"not null;type:text;comment:业务变量id"`
	Env        *Env      `json:"env" gorm:"constraint:onDelete:cascade"`
	Business   *Business `json:"business" gorm:"constraint:onDelete:cascade"`
}

type BusinessRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Business, error)
	GetByKeyword(ctx kratosx.Context, key string) (*Business, error)
	PageServerBusiness(ctx kratosx.Context, req *types.PageServerBusinessRequest) ([]*Business, uint32, error)
	AllServerBusiness(ctx kratosx.Context, id uint32) ([]*Business, error)
	Create(ctx kratosx.Context, c *Business) (uint32, error)
	Update(ctx kratosx.Context, c *Business) error
	Delete(ctx kratosx.Context, uint322 uint32) error
	GetValues(ctx kratosx.Context, bid uint32) ([]*BusinessValue, error)
	GetEnvValues(ctx kratosx.Context, eid, sid uint32) ([]*BusinessValue, error)
	UpdateValue(ctx kratosx.Context, rv *BusinessValue) error
}

type BusinessUseCase struct {
	config  *config.Config
	repo    BusinessRepo
	srvRepo ServerRepo
}

func NewBusinessUseCase(config *config.Config, repo BusinessRepo) *BusinessUseCase {
	return &BusinessUseCase{config: config, repo: repo}
}

// Get 获取指定业务变量信息
func (u *BusinessUseCase) Get(ctx kratosx.Context, id uint32) (*Business, error) {
	business, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return business, nil
}

// GetByKeyword 获取指定标识的业务变量信息
func (u *BusinessUseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*Business, error) {
	business, err := u.repo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return business, nil
}

// Add 添加业务变量信息
func (u *BusinessUseCase) Add(ctx kratosx.Context, business *Business) (uint32, error) {
	if business.Values != nil {
		for _, rv := range business.Values {
			if err := u.CheckBusinessValue(ctx, rv.BusinessID, rv.Value); err != nil {
				return 0, err
			}
		}
	}
	id, err := u.repo.Create(ctx, business)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新业务变量信息
func (u *BusinessUseCase) Update(ctx kratosx.Context, business *Business) error {
	if err := u.repo.Update(ctx, business); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除业务变量信息
func (u *BusinessUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// PageServerBusiness 获取指定服务的业务变量
func (u *BusinessUseCase) PageServerBusiness(ctx kratosx.Context, req *types.PageServerBusinessRequest) ([]*Business, uint32, error) {
	list, total, err := u.repo.PageServerBusiness(ctx, req)
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// AllBusinessValue 获取指定业务变量的所有环境值
func (u *BusinessUseCase) AllBusinessValue(ctx kratosx.Context, bid uint32) ([]*BusinessValue, error) {
	list, err := u.repo.GetValues(ctx, bid)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return list, nil
}

// AllBusinessField 获取指定服务业务变量的所有可用字段
func (u *BusinessUseCase) AllBusinessField(ctx kratosx.Context, sid uint32) ([]string, error) {
	list, err := u.repo.AllServerBusiness(ctx, sid)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	var arr []string
	for _, item := range list {
		arr = append(arr, item.Keyword)
	}
	return arr, nil
}

// UpdateBusinessValue 更新指定业务变量的值
func (u *BusinessUseCase) UpdateBusinessValue(ctx kratosx.Context, rv *BusinessValue) error {
	if err := u.CheckBusinessValue(ctx, rv.BusinessID, rv.Value); err != nil {
		return err
	}
	if err := u.repo.UpdateValue(ctx, rv); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// CheckBusinessValue 检测业务变量值是否合法
func (u *BusinessUseCase) CheckBusinessValue(ctx kratosx.Context, bid uint32, values string) error {
	return nil
}
