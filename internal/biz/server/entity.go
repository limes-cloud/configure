package server

type Server struct {
	Id          uint32  `json:"id"`
	Token       string  `json:"token"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      *bool   `json:"status"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}
