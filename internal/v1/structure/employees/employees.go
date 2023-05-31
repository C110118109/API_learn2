package employees

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Employee struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 員工編號
	EmployeeID string `gorm:"primaryKey;uuid_generate_v4();column:employee_id;type:uuid;" json:"employee_id,omitempty"`
	// 員工姓名
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 部門編號
	DepartmentID string `gorm:"column:department_id;type:uuid;" json:"department_id,omitempty"`
	// 角色編號
	RoleID string `gorm:"column:role_id;type:uuid;" json:"role_id,omitempty"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 員工編號
	EmployeeID string `json:"employee_id,omitempty"`
	// 員工姓名
	Name string `json:"name,omitempty"`
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"`
	// 角色編號
	RoleID string `json:"role_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 員工編號
	EmployeeID string `json:"employee_id,omitempty"`
	// 員工姓名
	Name string `json:"name,omitempty"`
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"`
	// 角色編號
	RoleID string `json:"role_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Created struct is used to create
type Created struct {
	// 員工姓名
	Name string `json:"name" binding:"required" validate:"required"`
	// 部門ID
	DepartmentID string `json:"department_id" binding:"required,uuid4" validate:"required"`
	// 角色編號
	RoleID string `json:"role_id" binding:"required,uuid4" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 員工編號
	EmployeeID string `json:"employee_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 部門ID
	DepartmentID *string `json:"department_id,omitempty" form:"department_id" binding:"omitempty,uuid4"`
	// 角色編號
	RoleID *string `json:"role_id,omitempty" form:"role_id" binding:"omitempty,uuid4"`
	// 員工姓名
	Name *string `json:"name,omitempty" form:"name"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Employees []*struct {
		// 員工編號
		EmployeeID string `json:"employee_id,omitempty"`
		// 部門ID
		DepartmentID string `json:"department_id,omitempty"`
		// 員工姓名
		Name string `json:"name,omitempty"`
		// 角色編號
		RoleID string `json:"role_id,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
	} `json:"employees"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 員工編號
	EmployeeID string `json:"employee_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 部門ID
	DepartmentID *string `json:"department_id,omitempty" binding:"omitempty,uuid4"`
	// 員工姓名
	Name *string `json:"name,omitempty"`
	// 角色編號
	RoleID *string `json:"role_id,omitempty" binding:"omitempty,uuid4"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "employees"
}
