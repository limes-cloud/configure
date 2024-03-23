package server

type PageServerRequest struct {
	Page        uint32   `json:"page"`
	PageSize    uint32   `json:"page_size"`
	ServerScope []string `json:"server_scope"`
	Keyword     *string  `json:"keyword"`
	IsBusiness  *bool    `json:"is_business"`
}
