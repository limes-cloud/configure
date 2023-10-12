package model

import "gorm.io/gorm"

type Scopes func(db *gorm.DB) *gorm.DB

type PageOptions struct {
	Page     int32
	PageSize int32
	Scopes   Scopes
}

type CreateModel struct {
	ID        int64 `json:"id"`
	CreatedAt int64 `json:"created_at,omitempty"`
}

type BaseModel struct {
	ID        int64 `json:"id"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

type DeleteModel struct {
	ID        int64          `json:"id"`
	CreatedAt int64          `json:"created_at,omitempty"`
	UpdatedAt int64          `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
