package server

import ktypes "github.com/limes-cloud/kratosx/types"

type Server struct {
	ktypes.BaseModel
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
