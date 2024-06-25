package template

type Template struct {
	Id          uint32 `json:"id"`
	ServerId    uint32 `json:"serverId"`
	Content     string `json:"content"`
	Version     string `json:"version"`
	IsUse       bool   `json:"isUse"`
	Format      string `json:"format"`
	Description string `json:"description"`
	Compare     string `json:"compare"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}
