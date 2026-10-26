package types

// SettingsItem 系统配置项（模板引用的业务变量 + 当前环境值）
type SettingsItem struct {
	Keyword     string `json:"keyword"`     // 变量 key（对应模板 ${keyword}）
	Type        string `json:"type"`        // 值类型：string | text | int | float | bool | enum | object | kv | array
	Description string `json:"description"` // 变量说明
	Extra       string `json:"extra"`       // 渲染元数据（JSON 字符串）
	Value       string `json:"value"`       // 当前值（原始字符串；未设置 = 空）
}

// GetSettingsReply 系统配置响应
type GetSettingsReply struct {
	App    string          `json:"app"`    // 应用标识
	EnvId  uint32          `json:"envId"`  // 环境 ID
	Format string          `json:"format"` // 当前模板格式（json | yaml）
	Items  []*SettingsItem `json:"items"`  // 配置项列表
}

// UpdateSettingsItem 系统配置保存项
type UpdateSettingsItem struct {
	Keyword string `json:"keyword"` // 变量 key
	Value   string `json:"value"`   // 配置值（原始字符串）
}