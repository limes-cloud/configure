package resource

type GetResourceRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type ListResourceRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
	Tag      *string `json:"tag"`
	Private  *bool   `json:"private"`
	ServerId *uint32 `json:"serverId"`
}
