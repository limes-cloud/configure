package types

type ListTemplateRequest struct {
	Page      uint32  `json:"page"`
	PageSize  uint32  `json:"pageSize"`
	App       string  `json:"app"`
	Namespace *string `json:"namespace"`
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
	EnvId     uint32 `json:"envId"`
	App       string `json:"app"`
	Namespace string `json:"namespace"`
	Format    string `json:"format"`
	Content   string `json:"content"`
}

type PreviewTemplateReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type PreviewCurrentTemplateRequest struct {
	EnvId     uint32 `json:"envId"`
	App       string `json:"app"`
	Namespace string `json:"namespace"`
}
