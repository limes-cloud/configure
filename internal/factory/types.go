package factory

type Compare struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Old  string `json:"old"`
	Cur  string `json:"cur"`
}

type ParseByContentRequest struct {
	EnvId    uint32 `json:"envId"`
	ServerId uint32 `json:"serverId"`
	Format   string `json:"format"`
	Content  string `json:"content"`
}

type TemplateValue struct {
	Value   string
	Exclude bool
}
