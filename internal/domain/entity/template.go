package entity

import "github.com/limes-cloud/kratosx/model"

// Template 配置模板（多版本管理，IsUse 标记当前生效版本）
type Template struct {
	model.BaseModel
	App         string `json:"app" gorm:"column:app"`                 // 所属应用标识
	Namespace   string `json:"namespace" gorm:"column:namespace"`     // 模板命名空间（空值按 conf 处理，见 service.normalizeTemplateNamespace）
	Content     string `json:"content" gorm:"column:content"`         // 模板内容（含 ${key} 占位符）
	Version     string `json:"version" gorm:"column:version"`         // 内容哈希（MD5 大写），同一内容复用同一版本
	IsUse       bool   `json:"isUse" gorm:"column:is_use"`            // 是否为当前生效版本（每应用+命名空间仅一条为 true）
	Format      string `json:"format" gorm:"column:format"`           // 内容格式：json | yaml
	Description string `json:"description" gorm:"column:description"` // 版本说明
	Compare     string `json:"compare" gorm:"column:compare"`         // 与上一版本的差异（JSON 数组），仅创建时生成
}
