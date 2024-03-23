package template

type PageTemplateRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
	ServerId uint32 `json:"server_id"`
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

type PreviewTemplateRequest struct {
	EnvId    uint32 `json:"env_id"`
	ServerId uint32 `json:"server_id"`
	Format   string `json:"format"`
	Content  string `json:"content"`
}

type PreviewTemplateReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type ParseTemplateRequest struct {
	EnvId    uint32 `json:"env_id"`
	ServerId uint32 `json:"server_id"`
}

type ParseTemplateReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}
