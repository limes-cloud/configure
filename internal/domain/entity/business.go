package entity

import (
	"strconv"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/configure/api/configure/errors"
)

type Business struct {
	ServerId    uint32  `json:"serverId" gorm:"column:server_id"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Type        string  `json:"type" gorm:"column:type"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type BusinessValue struct {
	EnvId      uint32    `json:"envId" gorm:"column:env_id"`
	BusinessId uint32    `json:"businessId" gorm:"column:business_id"`
	Value      string    `json:"value" gorm:"column:value"`
	Business   *Business `json:"business"`
	types.BaseModel
}

func (bv *BusinessValue) MarshalValue(tp string) (string, error) {
	var value = bv.Value
	switch tp {
	case "int":
		if _, err := strconv.Atoi(value); err != nil {
			return "", errors.BusinessValueTypeError()
		}
		return value, nil
	case "float":
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return "", errors.BusinessValueTypeError()
		}
		return value, nil
	case "string":
		return value, nil
	case "bool":
		if value != "true" && value != "false" {
			return "", errors.BusinessValueTypeError()
		}
		return value, nil
	case "object":
		var m any
		if err := json.Unmarshal([]byte(value), &m); err != nil {
			return "", errors.BusinessValueTypeError()
		}
		content, _ := json.Marshal(m)
		return string(content), nil
	default:
		return "", errors.BusinessValueTypeError()
	}
}
