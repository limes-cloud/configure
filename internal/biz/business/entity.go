package business

import ktypes "github.com/limes-cloud/kratosx/types"

type Business struct {
	ktypes.BaseModel
	ServerID    uint32 `json:"server_id"`
	Keyword     string `json:"keyword"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type BusinessValue struct {
	ktypes.BaseModel
	EnvId      uint32    `json:"env_id"`
	BusinessId uint32    `json:"business_id"`
	Value      string    `json:"value"`
	Business   *Business `json:"business" gorm:"foreignKey:business_id;references:id"`
}
