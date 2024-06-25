package env

type GetEnvRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type ListEnvRequest struct {
	Keyword *string `json:"keyword"`
	Name    *string `json:"name"`
	Status  *bool   `json:"status"`
}
