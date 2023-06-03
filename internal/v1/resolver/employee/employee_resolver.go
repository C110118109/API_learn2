package employee

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	//departmentModel "eirc.app/internal/v1/structure/departments"

	employeeModel "eirc.app/internal/v1/structure/employees"
	//roleModel "eirc.app/internal/v1/structure/roles"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *employeeModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	employee, err := r.EmployeeService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, employee.EmployeeID)
}

func (r *resolver) List(input *employeeModel.Fields) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	output := &employeeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, employees, err := r.EmployeeService.List(input)
	employeeByte, err := json.Marshal(employees)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(employeeByte, &output.Employees)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetBySingle(input *employeeModel.Base) interface{} {
	employee, err := r.EmployeeService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &employeeModel.Single{}
	employeeByte, _ := json.Marshal(employee)
	err = json.Unmarshal(employeeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.DepartmentName = employee.Departments.Name
	output.DepartmentCode = employee.Departments.DepartmentCode
	output.RoleName = employee.Roles.Name
	return code.GetCodeMessage(code.Successful, output)

}

func (r *resolver) GetByID(input *employeeModel.Field) interface{} {
	//input.IsDeleted = util.PointerBool(false)
	employee, err := r.EmployeeService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &employeeModel.Single{}
	employeeByte, _ := json.Marshal(employee)
	err = json.Unmarshal(employeeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *employeeModel.Updated) interface{} {
	_, err := r.EmployeeService.GetByID(&employeeModel.Field{EmployeeID: input.EmployeeID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.EmployeeService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *employeeModel.Updated) interface{} {
	employee, err := r.EmployeeService.GetByID(&employeeModel.Field{EmployeeID: input.EmployeeID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.EmployeeService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, employee.EmployeeID)
}
