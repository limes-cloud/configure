package entity

import "github.com/limes-cloud/kratosx/model"

type TemplateNamespace struct {
	App         string  `json:"app" gorm:"column:app"`
	Namespace   string  `json:"namespace" gorm:"column:namespace"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	Sort        uint32  `json:"sort" gorm:"column:sort"`
	model.BaseModel
}

func (TemplateNamespace) TableName() string {
	return "template_namespace"
}
