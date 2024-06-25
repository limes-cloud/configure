package business

type Business struct {
	Id          uint32  `json:"id"`
	ServerId    uint32  `json:"serverId"`
	Keyword     string  `json:"keyword"`
	Type        string  `json:"type"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}

type BusinessValue struct {
	Id         uint32    `json:"id"`
	EnvId      uint32    `json:"envId"`
	BusinessId uint32    `json:"businessId"`
	Value      string    `json:"value"`
	CreatedAt  int64     `json:"createdAt"`
	UpdatedAt  int64     `json:"updatedAt"`
	Business   *Business `json:"business"`
}
