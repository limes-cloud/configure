package types

type PageResourceRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Keyword  *string `json:"keyword"`
	Tag      *string `json:"tag"`
}

type PageServerResourceRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
	ServerID uint32 `json:"server_id"`
}
