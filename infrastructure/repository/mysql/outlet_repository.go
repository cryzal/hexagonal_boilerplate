package mysql

import (
	"context"
	"gorm.io/gorm"
	"hexagonal_boilerplate/core/entities"
	"hexagonal_boilerplate/infrastructure/repository/mysql/mapper"
	"hexagonal_boilerplate/infrastructure/repository/mysql/models"
	"strconv"
)

type (
	RepositoryOutlet struct {
		*gorm.DB
	}
)

func NewOutletRepo(db *gorm.DB) *RepositoryOutlet {
	return &RepositoryOutlet{db}
}

// model returns query model with context with or without transaction extracted from context
func (r *RepositoryOutlet) model(ctx context.Context) *gorm.DB {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.WithContext(ctx)
	}
	return r.WithContext(ctx)
}
func (r *RepositoryOutlet) Get(ID string) (*entities.Outlet, error) {
	var outletModel models.OutletModel
	query := r.DB.Where("code=?", ID).Take(&outletModel)
	if query.Error != nil {
		return nil, query.Error
	}

	entity, err := mapper.MapToEntities(&outletModel)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (r *RepositoryOutlet) Create(c context.Context, outlet *entities.Outlet) error {
	objModel := mapper.MapToWriteModels(outlet)
	store := r.model(c).Create(objModel)

	if err := store.Error; err != nil {
		return err
	}

	primaryID := strconv.FormatInt(*objModel.ID, 16)
	outlet.SetID(&primaryID, objModel.Code)
	return nil
}

func (r *RepositoryOutlet) Update(c context.Context, outlet *entities.Outlet) error {
	objModel := mapper.MapToWriteModels(outlet)

	update := r.model(c).Updates(objModel)
	if update.Error != nil {
		return update.Error
	}
	return nil
}
