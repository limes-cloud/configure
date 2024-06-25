package resource

type Resource struct {
	Id              uint32            `json:"id"`
	Keyword         string            `json:"keyword"`
	Fields          string            `json:"fields"`
	Tag             string            `json:"tag"`
	Private         *bool             `json:"private"`
	Description     *string           `json:"description"`
	CreatedAt       int64             `json:"createdAt"`
	UpdatedAt       int64             `json:"updatedAt"`
	Servers         []*Server         `json:"servers"`
	ResourceServers []*ResourceServer `json:"resourceServers"`
}

type Server struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type ResourceServer struct {
	ServerId   uint32 `json:"serverId"`
	ResourceId uint32 `json:"resourceId"`
}

type ResourceValue struct {
	Id         uint32    `json:"id"`
	EnvId      uint32    `json:"envId"`
	ResourceId uint32    `json:"resourceId"`
	Value      string    `json:"value"`
	CreatedAt  int64     `json:"createdAt"`
	UpdatedAt  int64     `json:"updatedAt"`
	Resource   *Resource `json:"resource"`
}
