package request

import (
	//"eirc.app/internal/v1/entity/equipment"
	//"eirc.app/internal/v1/presenter/employee"
	"eirc.app/internal/v1/service/department"
	"eirc.app/internal/v1/service/employee"
	"eirc.app/internal/v1/service/item"
	"eirc.app/internal/v1/service/request"
	"eirc.app/internal/v1/service/role"
	"eirc.app/internal/v1/service/request_itemlist"

	model "eirc.app/internal/v1/structure/requests"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	RequestService    request.Service
	DepartmentService department.Service
	EmployeeService   employee.Service
	RoleService  role.Service
	ItemService  item.Service
	RequestItemListService    request_itemlist.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		RequestService:    request.New(db),
		DepartmentService: department.New(db),
		EmployeeService:   employee.New(db),
		ItemService:  item.New(db),
		RoleService:  role.New(db),
		RequestItemListService:  request_itemlist.New(db),
	}
}
