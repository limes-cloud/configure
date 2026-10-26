package pkg

// ValidOrder 校验排序方向是否合法（白名单校验，纵深防御；正常入口已由 proto validate 拦截）
func ValidOrder(v string) bool {
	switch v {
	case "asc", "desc":
		return true
	}
	return false
}

// ValidOrderBy 校验排序列是否合法（与 proto validate 的 in 规则保持一致）
func ValidOrderBy(v string) bool {
	switch v {
	case "id":
		return true
	}
	return false
}
