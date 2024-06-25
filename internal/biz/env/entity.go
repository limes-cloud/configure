package env

type Env struct {
	Id          uint32  `json:"id"`
	Token       string  `json:"token"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Status      *bool   `json:"status"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}
