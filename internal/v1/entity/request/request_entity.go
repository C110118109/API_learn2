package request

import (
	"eirc.app/internal/pkg/log"
	model "eirc.app/internal/v1/structure/requests"
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
	if input.RequestID != "" {
		db.Where("re_id = ?", input.RequestID)
	}

	err = db.First(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("re_id = ?", input.RequestID)

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

func (e *entity) RequestDetailListUser(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Preload("Detail").
		Order("requests.created_time desc").
		Offset(int((input.Page - 1) * input.Limit)).Limit(int(input.Limit)).Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByReIDRequestDetailListUser(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Where("re_id = ?", input.RequestID).First(&output).Error

	return output, err
}
