package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Configure struct {
	types.BaseModel
	ServerID      uint32       `json:"server_id" gorm:"not null;comment:服务id"`
	EnvironmentID uint32       `json:"environment_id" gorm:"not null;comment:环境id"`
	Content       string       `json:"content" gorm:"not null;type:text;comment:配置内容"`
	Version       string       `json:"version" gorm:"not null;size:32;comment:配置版本"`
	Format        string       `json:"format" gorm:"not null;size:32;comment:配置格式"`
	Description   *string      `json:"description" gorm:"size:128;comment:配置描述"`
	Server        *Server      `json:"server" gorm:"constraint:onDelete:cascade"`
	Environment   *Environment `json:"environment" gorm:"constraint:onDelete:cascade"`
}

func (t *Configure) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(t).Error
}

func (t *Configure) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(t).Error
}

func (t *Configure) OneBySrvAndEnv(ctx kratosx.Context, srvId, envId uint32) error {
	return ctx.DB().First(t, "server_id=? and environment_id=?", srvId, envId).Error
}
