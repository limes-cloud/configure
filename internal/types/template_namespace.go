package types

type ListTemplateNamespaceRequest struct {
	Page      uint32  `json:"page"`
	PageSize  uint32  `json:"pageSize"`
	App       string  `json:"app"`
	Namespace *string `json:"namespace"`
	Name      *string `json:"name"`
}
