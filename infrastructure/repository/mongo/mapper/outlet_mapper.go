package mapper

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hexagonal_boilerplate/core/entities"
	"hexagonal_boilerplate/infrastructure/repository/mongo/models"
)

func MapToWriteModels(outlet *entities.Outlet) *models.OutletModel {
	model := models.OutletModel{
		Name:  outlet.Name,
		Email: outlet.Email,
		Phone: outlet.Phone,
	}
	if outlet.ID != nil {
		objID, err := primitive.ObjectIDFromHex(*outlet.ID)
		if err != nil {
			panic(err)
		}
		model.ID = objID
	}

	return &model
}

func MapToEntities(outletModel *models.OutletModel) (*entities.Outlet, error) {
	primitiveID := outletModel.ID
	primitiveIDString := primitiveID.Hex()
	objEntity := &entities.Outlet{
		ID:    &primitiveIDString,
		Code:  &primitiveIDString,
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
