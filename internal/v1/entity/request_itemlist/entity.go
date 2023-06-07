package request_itemlist

import (
	model "eirc.app/internal/v1/structure/request_itemlists"
	"gorm.io/gorm"
)

type Entity interface {
	WithTrx(tx *gorm.DB) Entity
	Created(input *model.Table) (err error)
	List(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByID(input *model.Field) (output *model.Table, err error)
	GetBySingle(input *model.Base) (output *model.Table, err error)
	ItemDetailUser(input *model.Fields) (amount int64, output []*model.Table, err error)
	GetByIIDItemDetailUser(input *model.Field) (output *model.Table, err error)
	//brfore : Deleted(input *model.Field) (err error)
	Deleted(input *model.Table) (err error)
	Updated(input *model.Table) (err error)
}

type entity struct {
	db *gorm.DB
}

func New(db *gorm.DB) Entity {
	return &entity{
		db: db,
	}
}

func (e *entity) WithTrx(tx *gorm.DB) Entity {
	return &entity{
		db: tx,
	}
}
