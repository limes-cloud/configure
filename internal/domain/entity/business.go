package entity

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx/model"

	"github.com/limes-cloud/configure/api/errors"
)

// BusinessExtra 业务变量渲染元数据（对应 business.extra JSON 列，字段可增量扩展）
type BusinessExtra struct {
	Options     []BusinessExtraOption `json:"options"`     // 枚举/多选供选项（label 可选，省略时展示 value）
	Min         *float64              `json:"min"`         // 数值下限（int/float）
	Max         *float64              `json:"max"`         // 数值上限（int/float）
	Unit        string                `json:"unit"`        // 单位后缀（如 MB）
	Placeholder string                `json:"placeholder"` // 输入占位提示
	Default     string                `json:"default"`     // 未设置值时的默认值（与 value 同格式）
	Required    bool                  `json:"required"`    // 保存时是否必填
	Rows        int                   `json:"rows"`        // text 行数
	MaxLength   int                   `json:"maxLength"`   // 字符串最大长度
}

// BusinessExtraOption 枚举选项
type BusinessExtraOption struct {
	Label string `json:"label"` // 展示文案（可空 = 使用 Value）
	Value string `json:"value"` // 存储值
}

// Business 业务配置项定义（按业务维度组织的配置项）
type Business struct {
	App         string  `json:"app" gorm:"column:app"`                 // 所属应用标识，如 manager / resource / gateway
	Keyword     string  `json:"keyword" gorm:"column:keyword"`         // 业务配置键（应用内唯一）
	Type        string  `json:"type" gorm:"column:type"`               // 值类型：string | text | int | float | bool | enum | object | kv | array
	Extra       string  `json:"extra" gorm:"column:extra"`             // 渲染元数据（JSON 字符串，见 ParseExtra）
	Description *string `json:"description" gorm:"column:description"` // 配置说明
	model.BaseModel         // Id(uint32 自增主键) + CreatedAt/UpdatedAt(int64 Unix 秒)
}

// BusinessValue 业务配置值（按 环境+业务配置项 维度存储，同一项每环境一条记录）
type BusinessValue struct {
	EnvId           uint32    `json:"envId" gorm:"column:env_id"`           // 所属环境 ID
	BusinessId      uint32    `json:"businessId" gorm:"column:business_id"` // 所属业务配置 ID
	Value           string    `json:"value" gorm:"column:value"`            // 配置值（按 Business.Type 校验格式）
	Business        *Business `json:"business"`                             // 关联的业务配置（Preload 加载）
	model.BaseModel           // Id(uint32 自增主键) + CreatedAt/UpdatedAt(int64 Unix 秒)
}

// ParseExtra 解析渲染元数据，解析失败返回 nil（不阻断业务，按无元数据处理）
func (b *Business) ParseExtra() *BusinessExtra {
	if b == nil || b.Extra == "" {
		return &BusinessExtra{}
	}
	extra := &BusinessExtra{}
	if err := jsoniter.Unmarshal([]byte(b.Extra), extra); err != nil {
		return &BusinessExtra{}
	}
	return extra
}

// ValidateExtra 校验渲染元数据合法性（仅允许合法 JSON 且字段类型正确）
func (b *Business) ValidateExtra() error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal([]byte(b.Extra), &BusinessExtra{})
}

// MarshalValue 按类型校验并序列化配置值，避免非法格式入库
// extra 为业务变量的渲染元数据（JSON 字符串），用于 int/float 范围与 enum/array 枚举校验，可传空
func (bv *BusinessValue) MarshalValue(tp string, extra string) (string, error) {
	var value = bv.Value
	extraObj := &BusinessExtra{}
	_ = jsoniter.Unmarshal([]byte(extra), extraObj)

	switch tp {
	case "int":
		n, err := strconv.Atoi(value)
		if err != nil {
			return "", errors.BusinessValueTypeError()
		}
		if extraObj.Min != nil && float64(n) < *extraObj.Min {
			return "", errors.BusinessValueTypeError()
		}
		if extraObj.Max != nil && float64(n) > *extraObj.Max {
			return "", errors.BusinessValueTypeError()
		}
		return value, nil
	case "float":
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return "", errors.BusinessValueTypeError()
		}
		if extraObj.Min != nil && f < *extraObj.Min {
			return "", errors.BusinessValueTypeError()
		}
		if extraObj.Max != nil && f > *extraObj.Max {
			return "", errors.BusinessValueTypeError()
		}
		return value, nil
	case "string", "text":
		return value, nil
	case "bool":
		if value != "true" && value != "false" {
			return "", errors.BusinessValueTypeError()
		}
		return value, nil
	case "enum":
		for _, opt := range extraObj.Options {
			if opt.Value == value {
				return value, nil
			}
		}
		return "", errors.BusinessValueTypeError()
	case "array":
		var arr []string
		if err := jsoniter.Unmarshal([]byte(value), &arr); err != nil {
			return "", errors.BusinessValueTypeError()
		}
		// 配置了合法枚举时校验元素在集合内（避免把任意值写入 array 类型变量）
		if len(extraObj.Options) > 0 {
			valid := map[string]bool{}
			for _, opt := range extraObj.Options {
				valid[opt.Value] = true
			}
			for _, v := range arr {
				if !valid[v] {
					return "", errors.BusinessValueTypeError()
				}
			}
		}
		// 统一序列化，保证存储格式为紧凑 JSON 数组
		content, _ := jsoniter.Marshal(arr)
		return string(content), nil
	case "object":
		var m any
		if err := jsoniter.Unmarshal([]byte(value), &m); err != nil {
			return "", errors.BusinessValueTypeError()
		}
		content, _ := jsoniter.Marshal(m)
		return string(content), nil
	case "kv":
		// 键值对：只支持单层（值必须是标量 string/number/bool/null，不允许嵌套对象/数组）
		var m map[string]any
		if err := jsoniter.Unmarshal([]byte(value), &m); err != nil {
			return "", errors.BusinessValueTypeError()
		}
		for _, v := range m {
			switch v.(type) {
			case string, float64, bool, nil:
			default:
				return "", errors.BusinessValueTypeError()
			}
		}
		content, _ := jsoniter.Marshal(m)
		return string(content), nil
	default:
		return "", errors.BusinessValueTypeError()
	}
}