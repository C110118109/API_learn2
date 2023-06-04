package request_itemlist

import (
	//"eirc.app/internal/v1/entity/equipment"
	//"eirc.app/internal/v1/presenter/employee"

	"eirc.app/internal/v1/service/request_itemlist"

	model "eirc.app/internal/v1/structure/request_itemlists"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	GetBySingle(input *model.Base) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	//RequestService         request.Service
	//ItemService            item.Service
	RequestItemListService request_itemlist.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		//RequestService:         request.New(db),
		//ItemService:            item.New(db),
		RequestItemListService: request_itemlist.New(db),
	}
}
