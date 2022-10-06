package mapper

import (
	"errors"
	"hexagonal_boilerplate/core/entities"
	"hexagonal_boilerplate/infrastructure/repository/mysql/models"
)

func MapToWriteModels(outlet *entities.Outlet) *models.OutletModel {
	model := models.OutletModel{
		Name:  outlet.Name,
		Email: outlet.Email,
		Phone: outlet.Phone,
	}
	if outlet.ID != nil {
		model.ID = outlet.ID
	}
	if outlet.Code != nil {
		model.Code = outlet.Code
	}
	return &model
}

func MapToEntities(outletModel *models.OutletModel) (*entities.Outlet, error) {
	objEntity := &entities.Outlet{
		ID:    outletModel.ID,
		Code:  outletModel.Code,
		Name:  outletModel.Name,
		Email: outletModel.Email,
		Phone: outletModel.Phone,
	}
	if objEntity.ID == nil {
		return nil, errors.New("primary key can't be null")
	}
	if objEntity.Code == nil {
		return nil, errors.New("id can't be null")
	}
	return objEntity, nil
}
