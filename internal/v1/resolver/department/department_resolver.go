package department

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	departmentModel "eirc.app/internal/v1/structure/departments"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *departmentModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	department, err := r.DepartmentService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, department.DepartmentID)
}

func (r *resolver) List(input *departmentModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &departmentModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, departments, err := r.DepartmentService.List(input)
	departmentsByte, err := json.Marshal(departments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(departmentsByte, &output.Departments)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *departmentModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	base, err := r.DepartmentService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontDepartment := &departmentModel.Single{}
	departmentsByte, _ := json.Marshal(base)
	err = json.Unmarshal(departmentsByte, &frontDepartment)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontDepartment)
}

func (r *resolver) Deleted(input *departmentModel.Updated) interface{} {
	// _, err := r.DepartmentService.GetByID(&departmentModel.Field{DepartmentID: input.DepartmentID,
	// 	IsDeleted: util.PointerBool(false)})
	_, err := r.DepartmentService.GetByID(&departmentModel.Field{DepartmentID: input.DepartmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.DepartmentService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *departmentModel.Updated) interface{} {
	// department, err := r.DepartmentService.GetByID(&departmentModel.Field{DepartmentID: input.DepartmentID,
	// 	IsDeleted: util.PointerBool(false)})
	department, err := r.DepartmentService.GetByID(&departmentModel.Field{DepartmentID: input.DepartmentID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.DepartmentService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, department.DepartmentID)
}
