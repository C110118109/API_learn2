package employees

import (
	"time"
	
	model "eirc.app/internal/v1/structure"
	"eirc.app/internal/v1/structure/departments"
	"eirc.app/internal/v1/structure/roles"
)

// Employee struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	// 員工編號
	EmployeeID string `gorm:"primaryKey;uuid_generate_v4();column:e_id;type:uuid;" json:"e_id,omitempty"`
	// 員工姓名
	Name string `gorm:"column:name;type:TEXT;" json:"name,omitempty"`
	// 部門編號
	DepartmentID string `gorm:"column:department_id;type:uuid;" json:"department_id,omitempty"`
	// department data
	Departments departments.Table `gorm:"foreignKey:DepartmentID;references:DepartmentID" json:"departments,omitempty"`
	// 角色編號
	RoleID string `gorm:"column:role_id;type:uuid;" json:"role_id,omitempty"`
	// role data
	Roles roles.Table `gorm:"foreignKey:RoleID;references:RoleID" json:"roles,omitempty"`
	// 創建時間
	CreatedTime time.Time `gorm:"column:created_time;type:TIMESTAMP;" json:"created_time"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 員工編號
	EmployeeID string `json:"e_id,omitempty"`
	// 員工姓名
	Name string `json:"name,omitempty"`
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"`
	// department data
	Departments departments.Base `json:"departments,omitempty"`
	// 角色編號
	RoleID string `json:"role_id,omitempty"`
	// role data
	Roles roles.Base `json:"roles,omitempty"`
	// 創建時間
	CreatedTime time.Time `json:"created_time"`
}

// Single return structure file
type Single struct {
	// 員工編號
	EmployeeID string `json:"e_id,omitempty"`
	// 員工姓名
	Name string `json:"name,omitempty"`
	// 部門編號
	DepartmentID string `json:"department_id,omitempty"`
	// 部門名稱
	DepartmentName string `json:"department_name,omitempty"`
	// 部門代號
	DepartmentCode string `json:"department_code,omitempty"`
	// 角色編號
	RoleID string `json:"role_id,omitempty"`
	// 角色名稱
	RoleName string `json:"role_name,omitempty"`
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
}

// Field is structure file for search
type Field struct {
	// 員工編號
	EmployeeID string `json:"e_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
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
		EmployeeID string `json:"e_id,omitempty"`
		// 部門ID
		DepartmentID string `json:"department_id,omitempty"`
		// 員工姓名
		Name string `json:"name,omitempty"`
		// 角色編號
		RoleID string `json:"role_id,omitempty"`
		// 創建時間
		CreatedTime time.Time `json:"created_time"`
	} `json:"employees"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 員工編號
	EmployeeID string `json:"e_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
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
