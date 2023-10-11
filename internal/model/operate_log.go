package model

import (
	"github.com/limes-cloud/kratos"
)

type OperateLog struct {
	ServerID      int64       `json:"server_id"`
	EnvironmentID int64       `json:"environment_id"`
	Config        string      `json:"config"`
	Compare       string      `json:"compare"`
	Operator      string      `json:"operator"`
	OperatorId    int64       `json:"operator_id"`
	Server        Server      `json:"server"`
	Environment   Environment `json:"environment"`
}

func (u *OperateLog) Create(ctx kratos.Context) error {
	return ctx.DB().Create(u).Error
}

func (u *OperateLog) Page(ctx kratos.Context, options PageOptions) ([]*OperateLog, error) {
	var list []*OperateLog

	db := ctx.DB().Model(u)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, db.Find(&list).Error
}

func (u *OperateLog) OneById(ctx kratos.Context, id int64) error {
	return ctx.DB().First(u, "id = ?", id).Error
}
