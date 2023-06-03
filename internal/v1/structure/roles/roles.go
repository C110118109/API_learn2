package roles

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	// 角色UUID
	RoleID string `gorm:"primaryKey;column:r_id;uuid_generate_v4()type:UUID;" json:"r_id,omitempty"`
	// 角色名稱
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 編號
	RoleID string `json:"r_id,omitempty"`
	// 角色名稱
	Name string `json:"name"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 編號
	RoleID string `json:"r_id,omitempty"`
	// 角色名稱
	Name string `json:"name,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 角色名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
}

// Field is structure file for search
type Field struct {
	// 編號
	RoleID string `json:"r_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 角色名稱
	Name *string `json:"name,omitempty" form:"name"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Roles []*struct {
		// 編號
		RoleID string `json:"r_id,omitempty"`
		// 角色名稱
		Name string `json:"name,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"roles"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 編號
	RoleID string `json:"r_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 角色名稱
	Name *string `json:"name,omitempty"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "roles"
}
