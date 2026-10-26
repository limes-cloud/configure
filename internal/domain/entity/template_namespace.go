package entity

import "github.com/limes-cloud/kratosx/model"

// TemplateNamespace 模板命名空间（用于按命名空间隔离管理模板）
type TemplateNamespace struct {
	App             string  `json:"app" gorm:"column:app"`                 // 所属应用标识
	Namespace       string  `json:"namespace" gorm:"column:namespace"`     // 命名空间标识（应用内唯一）
	Name            string  `json:"name" gorm:"column:name"`               // 命名空间名称
	Description     *string `json:"description" gorm:"column:description"` // 命名空间说明
	Sort            uint32  `json:"sort" gorm:"column:sort"`               // 排序权重（升序）
	model.BaseModel         // Id(uint32 自增主键) + CreatedAt/UpdatedAt(int64 Unix 秒)
}

func (TemplateNamespace) TableName() string {
	return "template_namespace"
}
