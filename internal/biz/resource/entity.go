package resource

import ktypes "github.com/limes-cloud/kratosx/types"

type Resource struct {
	ktypes.BaseModel
	Keyword         string            `json:"keyword"`
	Description     string            `json:"description"`
	Fields          string            `json:"fields"`
	Tag             string            `json:"tag"`
	Private         *bool             `json:"private"`
	ResourceServers []*ResourceServer `json:"resource_servers"`
}

type ResourceValue struct {
	ktypes.BaseModel
	EnvId      uint32    `json:"env_id"`
	ResourceId uint32    `json:"resource_id"`
	Value      string    `json:"value"`
	Resource   *Resource `json:"resource" gorm:"foreignKey:resource_id;references:id"`
}

type ResourceServer struct {
	ktypes.CreateModel
	ServerId   uint32 `json:"server_id"`
	ResourceId uint32 `json:"resource_id"`
}
