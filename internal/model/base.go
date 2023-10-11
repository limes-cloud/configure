package model

import "gorm.io/gorm"

type Scopes func(db *gorm.DB) *gorm.DB

type PageOptions struct {
	Page     int32
	PageSize int32
	Scopes   Scopes
}

type CreateModel struct {
	ID        int64 `json:"id" gorm:"primary_key;auto_increment;size:32;comment:主键ID"`
	CreatedAt int64 `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
}

type BaseModel struct {
	ID        int64 `json:"id" gorm:"primary_key;auto_increment;size:32;comment:主键ID"`
	CreatedAt int64 `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64 `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteModel struct {
	ID        int64          `json:"id" gorm:"primary_key;auto_increment;size:32;comment:主键ID"`
	CreatedAt int64          `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64          `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index;comment:删除时间"`
}
