package department

import (
	"eirc.app/internal/v1/service/department"
	model "eirc.app/internal/v1/structure/departments"
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
	DepartmentService department.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		DepartmentService: department.New(db),
	}
}
