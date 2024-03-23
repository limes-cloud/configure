package template

import ktypes "github.com/limes-cloud/kratosx/types"

type Template struct {
	ktypes.BaseModel
	ServerID    uint32 `json:"server_id"`
	Content     string `json:"content"`
	Version     string `json:"version"`
	IsUse       bool   `json:"is_use"`
	Format      string `json:"format"`
	Description string `json:"description"`
	Compare     string `json:"compare"`
}
