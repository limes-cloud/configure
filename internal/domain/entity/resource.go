package entity

import "github.com/limes-cloud/kratosx/model"

// Resource 资源配置定义（资源视图片，Fields 为资源字段的逗号分隔列表）
type Resource struct {
	Keyword         string         `json:"keyword" gorm:"column:keyword"`         // 资源标识（唯一），模板中通过 ${keyword.field} 引用
	Fields          string         `json:"fields" gorm:"column:fields"`           // 资源字段列表，逗号分隔（如 host,port）
	Tag             string         `json:"tag" gorm:"column:tag"`                 // 资源标签，用于分类检索
	Private         *bool          `json:"private" gorm:"column:private"`         // 私有资源：true=仅绑定 App 可见；false=全局公共；NULL=视为公共
	Description     *string        `json:"description" gorm:"column:description"` // 资源说明
	ResourceApps    []*ResourceApp `json:"resourceApps"`                          // 绑定的应用列表（关联查询，非表字段）
	model.BaseModel                // Id(uint32 自增主键) + CreatedAt/UpdatedAt(int64 Unix 秒)
}

// ResourceApp 资源配置与应用的绑定关系
type ResourceApp struct {
	App        string `json:"app" gorm:"column:app"`                // 绑定的应用标识
	ResourceId uint32 `json:"resourceId" gorm:"column:resource_id"` // 资源 ID
}

func (ResourceApp) TableName() string {
	return "resource_app"
}

// ResourceValue 资源配置值（按 环境+资源 维度存储，同一资源每环境一条记录）
type ResourceValue struct {
	EnvId           uint32    `json:"envId" gorm:"column:env_id"`           // 所属环境 ID
	ResourceId      uint32    `json:"resourceId" gorm:"column:resource_id"` // 所属资源 ID
	Value           string    `json:"value" gorm:"column:value"`            // 配置值（JSON 对象，键为 Resource.Fields 中定义的字段）
	Resource        *Resource `json:"resource"`                             // 关联的资源（Preload 加载）
	model.BaseModel           // Id(uint32 自增主键) + CreatedAt/UpdatedAt(int64 Unix 秒)
}
