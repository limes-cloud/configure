package pkg

import (
	"encoding/json"
	"strconv"
)

func CheckType(tp string, value string) bool {
	switch tp {
	case "int":
		_, err := strconv.Atoi(value)
		return err == nil
	case "float":
		_, err := strconv.ParseFloat(value, 64)
		return err == nil
	case "string":
		return true
	case "bool":
		return value == "true" || value == "false"
	case "object":
		m := map[string]any{}
		return json.Unmarshal([]byte(value), &m) == nil
	default:
		return false
	}
}
