package types

type PageServerBusinessRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	ServerID uint32  `json:"server_id"`
	Keyword  *string `json:"keyword"`
	Tag      *string `json:"tag"`
}
