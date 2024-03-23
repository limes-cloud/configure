package configure

import ktypes "github.com/limes-cloud/kratosx/types"

type Configure struct {
	ktypes.BaseModel
	ServerId    uint32  `json:"server_id"`
	EnvId       uint32  `json:"env_id"`
	Content     string  `json:"content"`
	Version     string  `json:"version"`
	Format      string  `json:"format"`
	Description *string `json:"description"`
}
