package requests

import (
	"time"

	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/employees"

	// "eirc.app/internal/v1/structure/items"
	"eirc.app/internal/v1/structure/request_itemlists"
)

// Table struct is database table struct
type Table struct {
	// 編號UUID
	RequestID string `gorm:"<-:create;primaryKey;column:re_id;uuid_generate_v4()type:UUID;" json:"re_id,omitempty"`

	//ItemLists request_itemlists.Table `gorm:"foreignkey:RequestID;references:re_id" json:"itemlists"`
	//申請人
	ApplicantID string `gorm:"column:applicant_id;type:UUID;" json:"applicant_id,omitempty"`
	//applicant data
	Employees employees.Table `gorm:"foreignKey:ApplicantID;references:e_id" json:"employees,omitempty"`
	// 請購產品列表ID
	//RequestItemListID string `gorm:"column:request_itemlist_id;type:UUID;" json:"request_itemlist_id,omitempty"`
	//request_itemlist data
	Detail     []request_itemlists.Table `gorm:"foreignKey:request_id;references:re_id" json:"detail"`
	ItemDetail []request_itemlists.Table `gorm:"foreignKey:item_id;references:re_id" json:"itemdetail"`
	// 請購是由
	Reason string `gorm:"column:reason;type:TEXT;" json:"reason,omitempty"`
	// 請購日期
	RequestDate string `gorm:"column:request_date;type:DATE;" json:"request_date"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 編號
	RequestID string `json:"re_id,omitempty"`

	Detail []request_itemlists.Base `json:"detail"`
	// 申請人
	ApplicantID string `json:"applicant_id,omitempty"`
	//applicant data
	Employees employees.Base `json:"employees,omitempty"`
	// 請購產品列表ID
	// RequestItemListID string `json:"request_itemlist_id,omitempty"`
	// request_itemlist data
	//ItemLists request_itemlists.Base `json:"itemlists,omitempty"`
	// 請購是由
	Reason string `json:"reason,omitempty"`
	// 請購日期
	RequestDate string `json:"request_date"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

type RequestDetail struct {
	Single
	Detail []request_itemlists.Base `json:"detail,omitempty"`
	// gorm:"foreignkey:request_id;references:re_id"

}

type AllRequestDetail struct {
	Request []*RequestDetail `json:"requests"`
	model.OutPage
}

// Single return structure file
type Single struct {
	// 編號
	RequestID string `json:"re_id,omitempty"`
	// 申請人
	ApplicantID string `json:"applicant_id,omitempty"`
	// 申請人姓名
	ApplicantName string `json:"applicant_name,omitempty"`
	// 請購產品列表ID
	//RequestItemListID string `json:"request_itemlist_id,omitempty"`
	// *用途
	//RequestItemListApplication string `json:"ri_application,omitempty"`
	// **產品名稱
	//RequestItemListName string `json:"ri_name,omitempty"`
	// **產品單位(規格)
	//RequestItemListUnit string `json:"ri_unit,omitempty"`
	// *請購數量
	//RequestItemListQuanity int64 `json:"ri_quanity,omitempty"`
	// **產品單價
	//RequestItemListPrice int64 `json:"ri_price,omitempty"`
	// *總價
	//RequestItemListTotal int64 `json:"ri_total,omitempty"`
	// 請購是由
	Reason string `json:"reason,omitempty"`
	// 請購日期
	RequestDate string `json:"request_date"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`

	// Detail []request_itemlists.Base `json:"detail,omitempty"`

	// Request []*Single `json:"requests"`
	// model.OutPage
}

// Created struct is used to create
type Created struct {
	// 申請人
	ApplicantID string `json:"applicant_id" binding:"required" validate:"required"`
	// 請購產品列表ID
	//RequestItemListID string `json:"request_itemlist_id" binding:"required,uuid4" validate:"required"`
	// 請購是由
	Reason string `json:"reason,omitempty" binding:"required" validate:"required"`
	// 請購日期
	RequestDate string `json:"request_date,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required"`
}

// Field is structure file for search
type Field struct {
	// 編號
	RequestID string `json:"re_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 申請人
	ApplicantID string `json:"applicant_id,omitempty" form:"applicant_id" binding:"omitempty,uuid4"`
	// 請購產品列表ID
	//RequestItemListID string `json:"request_itemlist_id,omitempty" form:"request_itemlist_id" binding:"omitempty,uuid4"`
	// 請購是由
	Reason string `json:"reason,omitempty" form:"reason"`
	// 請購日期
	RequestDate string `json:"request_date,omitempty" form:"request_date"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Requests []*struct {
		// 編號
		RequestID string `json:"re_id,omitempty"`
		// 申請人
		ApplicantID string `json:"applicant_id,omitempty"`
		// 請購產品列表ID
		//RequestItemListID string `json:"request_itemlist_id,omitempty"`
		// 請購是由
		Reason string `json:"reason,omitempty"`
		// 請購日期
		RequestDate string `json:"request_date,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"requests"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 編號
	RequestID string `json:"re_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 申請人
	ApplicantID string `json:"applicant_id,omitempty" binding:"omitempty,uuid4"`
	// 請購產品列表ID
	//RequestItemListID string `json:"request_itemlist_id,omitempty" binding:"omitempty,uuid4"`
	// 請購是由
	Reason string `json:"reason,omitempty"`
	// 請購日期
	RequestDate string `json:"request_date,omitempty"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "requests"
}
