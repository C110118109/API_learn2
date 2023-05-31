package item

import (
	"eirc.app/internal/v1/service/item"
	model "eirc.app/internal/v1/structure/items"
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
	ItemService item.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		ItemService: item.New(db),
	}
}
