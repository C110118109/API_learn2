package request

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	//employeeModel "eirc.app/internal/v1/structure/employees"
	//"eirc.app/internal/v1/entity/employee"
	requestModel "eirc.app/internal/v1/structure/requests"
	//requesitemlistModel "eirc.app/internal/v1/structure/request_itemlists"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *requestModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	request, err := r.RequestService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, request.RequestID)
}

func (r *resolver) List(input *requestModel.Fields) interface{} {

	output := &requestModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, requests, err := r.RequestService.List(input)
	requestsByte, err := json.Marshal(requests)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(requestsByte, &output.Requests)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetBySingle(input *requestModel.Base) interface{} {
	request, err := r.RequestService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &requestModel.Single{}
	requestByte, _ := json.Marshal(request)
	err = json.Unmarshal(requestByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.ApplicantName = request.Employees.Name

	// if applicantBase,err:= r.RequestService.GetBySingle(&requestModel.Field{
	// 	RequestID: *&requestBase.ApplicantID,
	// });err != nil{
	// 	output.ApplicantName=""
	// }else{
	// 	output.ApplicantName=*applicantBase.Name
	// }

	return code.GetCodeMessage(code.Successful, output)

}

func (r *resolver) GetByID(input *requestModel.Field) interface{} {

	request, err := r.RequestService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &requestModel.Single{}
	requestsByte, _ := json.Marshal(request)
	err = json.Unmarshal(requestsByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *requestModel.Updated) interface{} {
	_, err := r.RequestService.GetByID(&requestModel.Field{RequestID: input.RequestID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RequestService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *requestModel.Updated) interface{} {
	request, err := r.RequestService.GetByID(&requestModel.Field{RequestID: input.RequestID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RequestService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, request.RequestID)
}
