package entity

import "github.com/limes-cloud/kratosx/model"

type Template struct {
	model.BaseModel
	App         string `json:"app" gorm:"column:app"`
	Namespace   string `json:"namespace" gorm:"column:namespace"`
	Content     string `json:"content" gorm:"column:content"`
	Version     string `json:"version" gorm:"column:version"`
	IsUse       bool   `json:"isUse" gorm:"column:is_use"`
	Format      string `json:"format" gorm:"column:format"`
	Description string `json:"description" gorm:"column:description"`
	Compare     string `json:"compare" gorm:"column:compare"`
}
