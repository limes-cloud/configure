package types

type ListConfigureRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"pageSize"`
	App      string `json:"app"`
	EnvId    uint32 `json:"envId"`
}

type CompareConfigureRequest struct {
	App   string `json:"app"`
	EnvId uint32 `json:"envId"`
}

type CompareConfigureReply struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Old  string `json:"old"`
	Cur  string `json:"cur"`
}

type WatcherConfigRequest struct {
	Token string
	App   string
}

type WatcherConfigureReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type WatcherConfigReplyFunc func(*WatcherConfigureReply) error
