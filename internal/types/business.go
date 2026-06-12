package types

type ListBusinessRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	App      string  `json:"app"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
}

type ListBusinessValueRequest struct {
	App        *string `json:"app"`
	BusinessId *uint32 `json:"businessId"`
	EnvId      *uint32 `json:"envId"`
}
