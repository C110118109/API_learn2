package employee

import (
	"eirc.app/internal/pkg/log"
	model "eirc.app/internal/v1/structure/employees"
	"gorm.io/gorm/clause"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_time desc").Find(&output).Error
	return amount, output, err
}

func (e *entity) GetBySingle(input *model.Base) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.EmployeeID != "" {
		db.Where("e_id = ?", input.EmployeeID)
	}

	err = db.First(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("e_id = ?", input.EmployeeID)
	// if input.IsDeleted != nil {
	// 	db.Where("is_deleted = ?", input.IsDeleted)
	// }

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Save(&input).Error

	return err
}
