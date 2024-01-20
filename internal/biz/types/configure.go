package types

type CompareConfigureRequest struct {
	ServerID uint32 `json:"server_id"`
	EnvID    uint32 `json:"env_id"`
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
	EnvID    uint32 `json:"server"`
	ServerID uint32 `json:"token"`
}

type WatcherConfigureReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type WatcherConfigReplyFunc func(*WatcherConfigureReply) error
