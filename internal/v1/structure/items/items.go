package items

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Item struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 產品編號
	ItemID string `gorm:"primaryKey;uuid_generate_v4();column:i_id;type:UUID;" json:"i_id,omitempty"`
	// 產品名稱
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 產品單位
	Unit string `gorm:"column:unit;type:TEXT;" json:"unit,omitempty"`
	// 產品價格
	Price int64 `gorm:"column:price;type:INT4;" json:"price,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 產品編號
	ItemID string `json:"i_id,omitempty"`
	// 產品名稱
	Name string `json:"name,omitempty"`
	// 產品單位
	Unit string `json:"unit,omitempty"`
	// 產品價格
	Price int64 `json:"price,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 產品編號
	ItemID string `json:"i_id,omitempty"`
	// 產品名稱
	Name string `json:"name,omitempty"`
	// 產品單位
	Unit string `json:"unit,omitempty"`
	// 產品價格
	Price int64 `json:"price,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 產品名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 產品單位
	Unit string `json:"unit,omitempty" binding:"required" validate:"required"`
	// 產品價格
	Price int64 `json:"price,omitempty" binding:"required" validate:"required"`
}

// Updated struct is used to update
type Updated struct {
	// 產品編號
	ItemID string `json:"i_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 產品名稱
	Name string `json:"name,omitempty"`
	// 產品單位
	Unit string `json:"unit,omitempty"`
	// 產品價格
	Price int64 `json:"price,omitempty"`
}

// Field is structure file for search
type Field struct {
	// 產品編號
	ItemID string `json:"i_id,omitempty"  binding:"omitempty,uuid4" swaggerignore:"true"`
	// 產品名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 產品單位
	Unit *string `json:"unit,omitempty" form:"unit"`
	// 產品價格
	Price *int64 `json:"price,omitempty" form:"price"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Items []*struct {
		// 產品編號
		ItemID string `json:"i_id,omitempty"`
		// 產品名稱
		Name string `json:"name,omitempty"`
		// 產品單位
		Unit string `json:"unit,omitempty"`
		// 產品價格
		Price int64 `json:"price,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"equipments"`
	model.OutPage
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "items"
}
