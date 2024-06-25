package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Business struct {
	ServerId    uint32  `json:"serverId" gorm:"column:server_id"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Type        string  `json:"type" gorm:"column:type"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type BusinessValue struct {
	EnvId      uint32    `json:"envId" gorm:"column:env_id"`
	BusinessId uint32    `json:"businessId" gorm:"column:business_id"`
	Value      string    `json:"value" gorm:"column:value"`
	Business   *Business `json:"business"`
	types.BaseModel
}
