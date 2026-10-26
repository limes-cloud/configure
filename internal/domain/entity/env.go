package entity

import "github.com/limes-cloud/kratosx/model"

// Env 环境信息（dev/test/prod 等环境定义，含客户端拉取配置的鉴权 Token）
type Env struct {
	Token           string  `json:"token" gorm:"column:token"`             // 环境凭证，客户端 Watch 配置时用于鉴权（MD5 大写）
	Keyword         string  `json:"keyword" gorm:"column:keyword"`         // 环境标识（唯一，如 dev / test / prod）
	Name            string  `json:"name" gorm:"column:name"`               // 环境名称
	Status          *bool   `json:"status" gorm:"column:status"`           // 环境状态：true=启用 false=禁用；NULL=禁用
	Description     *string `json:"description" gorm:"column:description"` // 环境描述
	model.BaseModel         // Id(uint32 自增主键) + CreatedAt/UpdatedAt(int64 Unix 秒)
}
