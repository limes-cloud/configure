package configure

//	type UpdateConfigureRequest struct {
//		ServerId    uint32 `json:"server_id"`
//		EnvId       uint32 `json:"env_id"`
//		Description string `json:"description"`
//	}
type CompareConfigureRequest struct {
	ServerId uint32 `json:"server_id"`
	EnvId    uint32 `json:"env_id"`
	Format   string `json:"format"`
	Content  string `json:"content"`
}

type CompareConfigureReply struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Old  string `json:"old"`
	Cur  string `json:"cur"`
}

type WatcherConfigRequest struct {
	EnvId    uint32
	ServerId uint32
}

type WatcherConfigureReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type WatcherConfigReplyFunc func(*WatcherConfigureReply) error
