package model

type User struct {
	ID                uint32 `json:"id"`
	DepartmentID      uint32 `json:"department_id"`
	DepartmentKeyword string `json:"department_keyword"`
	RoleID            uint32 `json:"role_id"`
	RoleKeyword       string `json:"role_keyword"`
	Name              string `json:"name"`
	Email             string `json:"email"`
}
