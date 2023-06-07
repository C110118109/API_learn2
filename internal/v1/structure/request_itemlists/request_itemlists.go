package request_itemlists

import (
	"time"

	model "eirc.app/internal/v1/structure"
	//"eirc.app/internal/v1/structure/employees"
	"eirc.app/internal/v1/structure/items"
)

// Request_itemlist struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 請購產品列表編號
	RequestItemListID string `gorm:"<-:create;primaryKey;uuid_generate_v4();column:ri_id;type:uuid;" json:"ri_id,omitempty"`
	//
	//Detail []items.Table `gorm:"foreignKey:i_id;references:ri_id" json:"detail"`
	//
	RequestID string `gorm:"<-:create;column:request_id;type:UUID;not null;" json:"request_id,omitempty"`
	// 產品列表編號
	ItemID string `gorm:"column:item_id;type:uuid;" json:"item_id,omitempty"`
	//item data
	Items []items.Table `gorm:"foreignKey:ItemID;references:ItemID" json:"items,omitempty"`
	// 請購單編號
	//RequestID string `gorm:"column:request_id;type:uuid;" json:"request_id,omitempty"`
	// 品名
	//Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 用途
	Application string `gorm:"column:application;type:TEXT;" json:"application,omitempty"`
	// 請購數量
	Quanity int64 `gorm:"column:quanity;type:INT4;" json:"quanity,omitempty"`
	// 產品單位(規格)
	//Unit string `gorm:"column:unit;type:TEXT;" json:"unit,omitempty"`
	// 單價
	//Price int64 `gorm:"column:price;type:INT4;" json:"price,omitempty"`
	// 總價
	Total int64 `gorm:"column:total;type:INT4;" json:"total,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 請購產品列表編號
	RequestItemListID string `json:"ri_id,omitempty"`

	//Detail []items.Base `json:"detail"`
	//
	RequestID string `json:"request_id,omitempty"`
	// 產品列表編號
	ItemID string `json:"item_id,omitempty"`
	// item data
	Items []items.Base `json:"items,omitempty"`
	// 品名
	//Name string `json:"name,omitempty"`
	// 請購單編號
	//RequestID string `json:"request_id,omitempty"`
	// 用途
	Application string `json:"application,omitempty"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty"`
	// 產品單位(規格)
	//Unit string `json:"unit,omitempty"`
	// 單價
	//Price int64 `json:"price,omitempty"`
	// 總價
	Total int64 `json:"total,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 請購產品列表編號
	RequestItemListID string `json:"ri_id,omitempty"`
	//
	RequestID string `json:"request_id,omitempty"`
	// 請購單編號
	//RequestID string `json:"request_id,omitempty"`
	// 用途
	Application string `json:"application,omitempty"`
	// 產品列表編號
	ItemID string `json:"item_id,omitempty"`
	// 產品名稱
	ItemName string `json:"item_name,omitempty"`
	// 產品單位(規格)
	ItemUnit string `json:"item_unit,omitempty"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty"`
	// 產品單價
	ItemPrice int64 `json:"item_price,omitempty"`
	// 總價
	Total int64 `json:"total,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建者姓名
	EmployeeName string `json:"employee_name,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	//
	RequestID string `json:"request_id,omitempty" binding:"required,uuid4" validate:"required"`
	// 品名
	//Name string `json:"name" binding:"required" validate:"required"`
	// 請購單編號
	//RequestID string `json:"request_id" binding:"required,uuid4" validate:"required"`
	// 產品列表編號
	ItemID string `json:"item_id" binding:"required,uuid4" validate:"required"`
	// 用途
	Application string `json:"application,omitempty" binding:"required" validate:"required"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty" binding:"required" validate:"required"`
	// 產品單位(規格)
	//Unit string `json:"unit,omitempty" binding:"required" validate:"required"`
	// 單價
	//Price int64 `json:"price,omitempty" binding:"required" validate:"required"`
	// 總價
	Total int64 `json:"total,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 請購產品列表編號
	RequestItemListID string `json:"ri_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 請購單編號
	//RequestID string `json:"request_id,omitempty" form:"request_id" binding:"omitempty,uuid4"`
	// 產品列表編號
	ItemID *string `json:"item_id,omitempty" form:"item_id" binding:"omitempty,uuid4"`
	// 用途
	Application string `json:"application,omitempty" form:"application"`
	// 品名
	//Name string `json:"name,omitempty" form:"name"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty" form:"quanity"`
	// 產品單位(規格)
	//Unit string `json:"unit,omitempty" form:"unit"`
	// 單價
	//Price int64 `json:"price,omitempty" form:"price"`
	// 總價
	Total int64 `json:"total,omitempty" form:"total"`
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
		RequestItemListID string `json:"ri_id,omitempty"`
		//
		RequestID string `json:"request_id,omitempty"`
		// 請購單編號
		//RequestID string `json:"request_id,omitempty"`
		// 產品列表編號
		ItemID string `json:"item_id,omitempty"`
		// 品名
		//Name string `json:"name,omitempty"`
		// 用途
		Application string `json:"application,omitempty"`
		// 請購數量
		Quanity int64 `json:"quanity,omitempty"`
		// 產品單位(規格)
		//Unit string `json:"unit,omitempty"`
		// 單價
		//Price int64 `json:"price,omitempty"`
		// 總價
		Total int64 `json:"total,omitempty"`
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
	RequestItemListID string `json:"ri_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 請購單編號
	//RequestID string `json:"request_id,omitempty" binding:"omitempty,uuid4"`
	// 品名
	//Name string `json:"name,omitempty"`
	// 產品列表編號
	ItemID string `json:"item_id,omitempty" binding:"omitempty,uuid4"`
	// 用途
	Application string `json:"application,omitempty"`
	// 請購數量
	Quanity int64 `json:"quanity,omitempty"`
	// 產品單位(規格)
	//Unit string `json:"unit,omitempty"`
	// 單價
	//Price int64 `json:"price,omitempty"`
	// 總價
	Total int64 `json:"total,omitempty"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "request_itemlists"
}

type ItemDetail struct {
	Single
	Detail []items.Base `json:"detail,omitempty"`
	// gorm:"foreignkey:request_id;references:re_id"

}

type AllItemDetail struct {
	Item []*ItemDetail `json:"item"`
	model.OutPage
}
