package request_itemlist

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	requesitemlistModel "eirc.app/internal/v1/structure/request_itemlists"

	//requestModel "eirc.app/internal/v1/structure/requests"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *requesitemlistModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	request_itemlist, err := r.RequestItemListService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, request_itemlist.RequestItemListID)
}

func (r *resolver) List(input *requesitemlistModel.Fields) interface{} {

	output := &requesitemlistModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, request_itemlists, err := r.RequestItemListService.List(input)
	request_itemlistsByte, err := json.Marshal(request_itemlists)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(request_itemlistsByte, &output.Request_itemlists)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetBySingle(input *requesitemlistModel.Base) interface{} {
	requesitemlist, err := r.RequestItemListService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &requesitemlistModel.Single{}
	requesitemlistByte, _ := json.Marshal(requesitemlist)
	err = json.Unmarshal(requesitemlistByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.ItemName = requesitemlist.Items.Name
	output.ItemUnit = requesitemlist.Items.Unit
	output.ItemPrice = requesitemlist.Items.Price
	return code.GetCodeMessage(code.Successful, output)

}

func (r *resolver) GetByID(input *requesitemlistModel.Field) interface{} {

	request_itemlist, err := r.RequestItemListService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &requesitemlistModel.Single{}
	request_itemlistsByte, _ := json.Marshal(request_itemlist)
	err = json.Unmarshal(request_itemlistsByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *requesitemlistModel.Updated) interface{} {
	_, err := r.RequestItemListService.GetByID(&requesitemlistModel.Field{RequestItemListID: input.RequestItemListID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RequestItemListService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *requesitemlistModel.Updated) interface{} {
	request_itemlist, err := r.RequestItemListService.GetByID(&requesitemlistModel.Field{RequestItemListID: input.RequestItemListID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RequestItemListService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, request_itemlist.RequestItemListID)
}
