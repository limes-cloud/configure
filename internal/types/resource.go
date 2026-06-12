package types

type ListResourceRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
	Tag      *string `json:"tag"`
	Private  *bool   `json:"private"`
	App      *string `json:"app"`
}

type ListResourceValueRequest struct {
	App        *string `json:"app"`
	ResourceId *uint32 `json:"resourceId"`
	EnvId      *uint32 `json:"envId"`
}
