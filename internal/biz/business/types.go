package business

type GetBusinessRequest struct {
	Id *uint32 `json:"id"`
}

type ListBusinessRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	ServerId *uint32 `json:"serverId"`
	Keyword  *string `json:"keyword"`
}
