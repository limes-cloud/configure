package server

type GetServerRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type ListServerRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
	Status   *bool   `json:"status"`
}
