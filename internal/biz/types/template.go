package types

type PageTemplateRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
	ServerID uint32 `json:"server_id"`
}

type CompareTemplateRequest struct {
	Id      uint32 `json:"id"`
	Format  string `json:"format"`
	Content string `json:"content"`
}

type CompareTemplateReply struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Old  string `json:"old"`
	Cur  string `json:"cur"`
}

type ParseByContentRequest struct {
	EnvId    uint32 `json:"env_id"`
	ServerId uint32 `json:"server_id"`
	Format   string `json:"format"`
	Content  string `json:"content"`
}

type TemplateValue struct {
	Value   string
	Exclude bool
}

type ParseRequest struct {
	EnvId    uint32 `json:"env_id"`
	ServerId uint32 `json:"server_id"`
}

type ParseReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}
