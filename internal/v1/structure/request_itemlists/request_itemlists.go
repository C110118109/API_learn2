package request_itemlists

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Request_itemlist struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 請購產品列表編號
	RequestItemListID string `gorm:"primaryKey;uuid_generate_v4();column:request_itemlist_id;type:uuid;" json:"request_itemlist_id,omitempty"`
	// 品名
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 用途
	Application string `gorm:"column:application;type:TEXT;" json:"application,omitempty"`
	// 請購單編號
	RequestID string `gorm:"column:request_id;type:uuid;" json:"request_id,omitempty"`
	// 請購數量
	Quanity int64 `gorm:"column:quanity;type:INT4;" json:"quanity,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 請購產品列表編號
	RequestItemListID string `json:"request_itemlist_id,omitempty"`
	// 品名
	Name string `json:"name,omitempty"`
	// 請購單編號
	RequestID string `json:"request_id,omitempty"`
	// 用途
	Application string `json:"application,omitempty"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 請購產品列表編號
	RequestItemListID string `json:"request_itemlist_id,omitempty"`
	// 品名
	Name string `json:"name,omitempty"`
	// 請購單編號
	RequestID string `json:"request_id,omitempty"`
	// 用途
	Application string `json:"application,omitempty"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 品名
	Name string `json:"name" binding:"required" validate:"required"`
	// 請購單編號
	RequestID string `json:"request_id" binding:"required,uuid4" validate:"required"`
	// 用途
	Application string `json:"application,omitempty" binding:"required" validate:"required"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 請購產品列表編號
	RequestItemListID string `json:"request_itemlist_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 請購單編號
	RequestID *string `json:"request_id,omitempty" form:"request_id" binding:"omitempty,uuid4"`
	// 用途
	Application *string `json:"application,omitempty" form:"application"`
	// 品名
	Name *string `json:"name,omitempty" form:"name"`
	// 請購數量
	Quanity *int64 `json:"quanity,omitempty" form:"quanity"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Request_itemlists []*struct {
		// 請購產品列表編號
		RequestItemListID string `json:"request_itemlist_id,omitempty"`
		// 請購單編號
		RequestID string `json:"request_id,omitempty"`
		// 品名
		Name string `json:"name,omitempty"`
		// 用途
		Application string `json:"application,omitempty"`
		// 請購數量
		Quanity int64 `json:"quanity,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"request_itemlists"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 請購產品列表編號
	RequestItemListID string `json:"request_itemlist_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 請購單編號
	RequestID *string `json:"request_id,omitempty" binding:"omitempty,uuid4"`
	// 品名
	Name *string `json:"name,omitempty"`
	// 用途
	Application string `json:"application,omitempty"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "request_itemlists"
}