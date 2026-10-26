package entity

import "github.com/limes-cloud/kratosx/model"

// Configure 应用在某环境下的最终配置（由当前模板 + 业务/资源值渲染生成）
type Configure struct {
	model.BaseModel
	App         string  `json:"app" gorm:"column:app"`                 // 所属应用标识
	EnvId       uint32  `json:"envId" gorm:"column:env_id"`            // 所属环境 ID
	Content     string  `json:"content" gorm:"column:content"`         // 渲染后的配置内容
	Version     string  `json:"version" gorm:"column:version"`         // 内容哈希（MD5 大写），用于版本比对与去重
	Format      string  `json:"format" gorm:"column:format"`           // 内容格式：json | yaml
	Description *string `json:"description" gorm:"column:description"` // 配置说明
}
