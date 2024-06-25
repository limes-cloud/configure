package configure

import ktypes "github.com/limes-cloud/kratosx/types"

type Configure struct {
	ktypes.BaseModel
	ServerId    uint32  `json:"serverId"`
	EnvId       uint32  `json:"envId"`
	Content     string  `json:"content"`
	Version     string  `json:"version"`
	Format      string  `json:"format"`
	Description *string `json:"description"`
}
