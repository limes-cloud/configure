package model

import (
	"github.com/limes-cloud/kratos"
)

type Configure struct {
	BaseModel
	ServerID      int64   `json:"server_id"`
	EnvironmentID int64   `json:"environment_id"`
	Content       string  `json:"content"`
	Version       string  `json:"version"`
	Format        string  `json:"format"`
	Description   *string `json:"description"`
	Operator      string  `json:"operator"`
	OperatorID    int64   `json:"operator_id"`
}

func (t *Configure) Create(ctx kratos.Context) error {
	return ctx.DB().Create(t).Error
}

func (t *Configure) Update(ctx kratos.Context) error {
	return ctx.DB().Updates(t).Error
}

func (t *Configure) OneBySrvAndEnv(ctx kratos.Context, srvId, envId int64) error {
	return ctx.DB().First(t, "server_id=? and environment_id=?", srvId, envId).Error
}
