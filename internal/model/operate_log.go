package model

import (
	"github.com/limes-cloud/kratos"
)

type OperateLog struct {
	ServerID      int64       `json:"server_id"`
	EnvironmentID int64       `json:"environment_id"`
	Config        string      `json:"config"`
	Compare       string      `json:"compare"`
	Server        Server      `json:"server"`
	Environment   Environment `json:"environment"`
}

func (ol *OperateLog) Create(ctx kratos.Context) error {
	return ctx.DB().Create(ol).Error
}

func (ol *OperateLog) Page(ctx kratos.Context, options PageOptions) ([]*OperateLog, int64, error) {
	var list []*OperateLog
	total := int64(0)

	db := ctx.DB().Model(ol)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, total, db.Find(&list).Error
}

func (ol *OperateLog) OneById(ctx kratos.Context, id int64) error {
	return ctx.DB().First(ol, "id = ?", id).Error
}
