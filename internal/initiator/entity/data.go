package entity

import managerEntity "github.com/limes-cloud/manager/api/entity"

const appKeyword = "configure"

var registerTables = map[string]struct{}{
	"env":                {},
	"business":           {},
	"business_value":     {},
	"resource":           {},
	"resource_app":       {},
	"resource_value":     {},
	"template":           {},
	"template_namespace": {},
	"configure":          {},
}

var ignoreFields = map[string]struct{}{
	"deleted_at": {},
}

func shouldRegisterTable(name string) bool {
	_, ok := registerTables[name]
	return ok
}

func shouldRegisterField(name string) bool {
	_, ok := ignoreFields[name]
	return !ok
}

func registerField(name string, comment string, index uint32) *managerEntity.RegisterEntityField {
	return &managerEntity.RegisterEntityField{
		Name:    name,
		Comment: comment,
		Index:   index,
	}
}
