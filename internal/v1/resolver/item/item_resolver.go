package item

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	itemModel "eirc.app/internal/v1/structure/items"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *itemModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	item, err := r.ItemService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, item.ItemID)
}

func (r *resolver) List(input *itemModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &itemModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, items, err := r.ItemService.List(input)
	itemsByte, err := json.Marshal(items)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(itemsByte, &output.Items)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *itemModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	base, err := r.ItemService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontItem := &itemModel.Single{}
	itemsByte, _ := json.Marshal(base)
	err = json.Unmarshal(itemsByte, &frontItem)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontItem)
}

func (r *resolver) Deleted(input *itemModel.Updated) interface{} {
	_, err := r.ItemService.GetByID(&itemModel.Field{ItemID: input.ItemID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ItemService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *itemModel.Updated) interface{} {
	item, err := r.ItemService.GetByID(&itemModel.Field{ItemID: input.ItemID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ItemService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, item.ItemID)
}
