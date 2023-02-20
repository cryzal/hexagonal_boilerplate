package outlet

import (
	"context"
	"hexagonal_boilerplate/core/domain/entities"
)

type (
	Repository interface {
		Get(ID string) (*entities.Outlet, error)
		Create(c context.Context, outlet *entities.Outlet) error
		Update(c context.Context, outlet *entities.Outlet) error
	}
)
