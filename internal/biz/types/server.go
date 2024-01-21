package types

type PageServerRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"page_size"`
	Keyword    *string `json:"keyword"`
	IsBusiness *bool   `json:"is_business"`
}
