package requests

import (
	"time"

	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/employees"
)

// Table struct is database table struct
type Table struct {
	// 編號UUID
	RequestID string `gorm:"primaryKey;column:re_id;uuid_generate_v4()type:UUID;" json:"re_id,omitempty"`
	//申請人
	ApplicantID string `gorm:"column:applicant;type:UUID;" json:"applicant,omitempty"`
	//applicant data
	Employees employees.Table `gorm:"foreignKey:ApplicantID;references:ApplicantID" json:"employees,omitempty"`
	// 請購產品列表ID
	RequestItemListID string `gorm:"column:request_itemlist_id;type:UUID;" json:"request_itemlist_id,omitempty"`
	// 請購是由
	Reason string `gorm:"column:reason;type:TEXT;" json:"reason,omitempty"`
	// 請購日期
	RequestDate time.Time `gorm:"column:request_date;type:TIMESTAMP;" json:"request_date"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 編號
	RequestID string `json:"re_id,omitempty"`
	// 申請人
	ApplicantID string `json:"applicant,omitempty"`
	//applicant data
	Employees employees.Base `json:"employees,omitempty"`
	// 請購產品列表ID
	RequestItemListID string `json:"request_itemlist_id,omitempty"`
	// 請購是由
	Reason string `json:"reason,omitempty"`
	// 請購日期
	RequestDate time.Time `json:"request_date"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 編號
	RequestID string `json:"re_id,omitempty"`
	// 申請人
	ApplicantID string `json:"applicant,omitempty"`
	// 申請人姓名
	ApplicantName string `json:"applicant_name,omitempty"`
	// 請購產品列表ID
	RequestItemListID string `json:"request_itemlist_id,omitempty"`
	// 請購是由
	Reason string `json:"reason,omitempty"`
	// 請購日期
	RequestDate time.Time `json:"request_date"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 申請人
	ApplicantID string `json:"applicant" binding:"required" validate:"required"`
	// 請購產品列表ID
	RequestItemListID string `json:"request_itemlist_id" binding:"required,uuid4" validate:"required"`
	// 請購是由
	Reason string `json:"reason,omitempty" binding:"required" validate:"required"`
	// 請購日期
	RequestDate time.Time `json:"request_date,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 編號
	RequestID string `json:"re_id,omitempty" binding:"omitempty" swaggerignore:"true"`
	// 申請人
	ApplicantID string `json:"applicant,omitempty" form:"applicant" binding:"omitempty,uuid4"`
	// 請購產品列表ID
	RequestItemListID string `json:"request_itemlist_id,omitempty" form:"request_itemlist_id" binding:"omitempty,uuid4"`
	// 請購是由
	Reason string `json:"reason,omitempty" form:"reason"`
	// 請購日期
	RequestDate time.Time `json:"request_date,omitempty" form:"request_date"`
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
		ApplicantID string `json:"applicant,omitempty"`
		// 請購產品列表ID
		RequestItemListID string `json:"request_itemlist_id,omitempty"`
		// 請購是由
		Reason string `json:"reason,omitempty"`
		// 請購日期
		RequestDate time.Time `json:"request_date,omitempty"`
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
	ApplicantID string `json:"applicant,omitempty" binding:"omitempty,uuid4"`
	// 請購產品列表ID
	RequestItemListID string `json:"request_itemlist_id,omitempty" binding:"omitempty,uuid4"`
	// 請購是由
	Reason string `json:"reason,omitempty"`
	// 請購日期
	RequestDate time.Time `json:"request_date,omitempty"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "requests"
}
