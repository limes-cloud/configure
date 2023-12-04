package model

import "github.com/limes-cloud/kratosx"

type Configure struct {
	BaseModel
	ServerID      uint32  `json:"server_id"`
	EnvironmentID uint32  `json:"environment_id"`
	Content       string  `json:"content"`
	Version       string  `json:"version"`
	Format        string  `json:"format"`
	Description   *string `json:"description"`
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
