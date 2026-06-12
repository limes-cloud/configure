package entity

import "github.com/limes-cloud/kratosx/model"

type Resource struct {
	Keyword      string         `json:"keyword" gorm:"column:keyword"`
	Fields       string         `json:"fields" gorm:"column:fields"`
	Tag          string         `json:"tag" gorm:"column:tag"`
	Private      *bool          `json:"private" gorm:"column:private"`
	Description  *string        `json:"description" gorm:"column:description"`
	ResourceApps []*ResourceApp `json:"resourceApps"`
	model.BaseModel
}

type ResourceApp struct {
	App        string `json:"app"  gorm:"column:app"`
	ResourceId uint32 `json:"resourceId"  gorm:"column:resource_id"`
}

func (ResourceApp) TableName() string {
	return "resource_app"
}

type ResourceValue struct {
	EnvId      uint32    `json:"envId" gorm:"column:env_id"`
	ResourceId uint32    `json:"resourceId" gorm:"column:resource_id"`
	Value      string    `json:"value" gorm:"column:value"`
	Resource   *Resource `json:"resource"`
	model.BaseModel
}
