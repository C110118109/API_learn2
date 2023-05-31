package request_itemlist

import (
	model "eirc.app/internal/v1/structure/request_itemlists"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	if input.RequestID != nil {
		db.Where("request_id = ?", input.RequestID)
	}

	// if input.Price != nil {
	// 	db.Where("price = ?", input.Price)
	// }

	// if input.Quanity != nil {
	// 	db.Where("name like %?%", *input.Quanity)
	// }

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("request_itemlist_id = ?", input.RequestItemListID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Save(&input).Error

	return err
}
