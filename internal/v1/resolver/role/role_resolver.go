package role

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	roleModel "eirc.app/internal/v1/structure/roles"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *roleModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	role, err := r.RoleService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, role.RoleID)
}

func (r *resolver) List(input *roleModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &roleModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, roles, err := r.RoleService.List(input)
	rolesByte, err := json.Marshal(roles)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(rolesByte, &output.Roles)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *roleModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	base, err := r.RoleService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontRole := &roleModel.Single{}
	rolesByte, _ := json.Marshal(base)
	err = json.Unmarshal(rolesByte, &frontRole)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontRole)
}

func (r *resolver) Deleted(input *roleModel.Updated) interface{} {
	// _, err := r.DepartmentService.GetByID(&departmentModel.Field{DepartmentID: input.DepartmentID,
	// 	IsDeleted: util.PointerBool(false)})
	_, err := r.RoleService.GetByID(&roleModel.Field{RoleID: input.RoleID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RoleService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *roleModel.Updated) interface{} {
	// department, err := r.DepartmentService.GetByID(&departmentModel.Field{DepartmentID: input.DepartmentID,
	// 	IsDeleted: util.PointerBool(false)})
	role, err := r.RoleService.GetByID(&roleModel.Field{RoleID: input.RoleID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RoleService.Updated(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, role.RoleID)
}
