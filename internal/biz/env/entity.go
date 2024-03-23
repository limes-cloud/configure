package env

import "github.com/limes-cloud/kratosx/types"

type Env struct {
	types.BaseModel
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Token       string `json:"token,omitempty"`
	Status      *bool  `json:"status"`
}
