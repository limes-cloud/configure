package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type OperateLog struct {
	types.BaseModel
	ServerID      uint32       `json:"server_id" gorm:"not null;comment:服务id"`
	EnvironmentID uint32       `json:"environment_id" gorm:"not null;comment:环境id"`
	Config        string       `json:"config" gorm:"not null;type:text;comment:配置内容"`
	Compare       string       `json:"compare" gorm:"not null;type:text;comment:变更内容"`
	Server        *Server      `json:"server" gorm:"constraint:onDelete:cascade"`
	Environment   *Environment `json:"environment" gorm:"constraint:onDelete:cascade"`
}

func (ol *OperateLog) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(ol).Error
}

func (ol *OperateLog) Page(ctx kratosx.Context, options types.PageOptions) ([]*OperateLog, uint32, error) {
	var list []*OperateLog
	total := int64(0)

	db := ctx.DB().Model(ol)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (ol *OperateLog) OneById(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(ol, "id = ?", id).Error
}
