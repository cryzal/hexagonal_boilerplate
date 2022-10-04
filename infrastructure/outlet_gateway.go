package infrastructure

import (
	"hexagonal_boilerplate/shared/config"

	"hexagonal_boilerplate/infrastructure/repository/mysql"
)

type OutletGateway struct {
	OutletRepo      *mysql.RepositoryOutlet
	TransactionRepo *mysql.RepositoryDbtransaction
}

func NewOutletGateway(cfg *config.Config) *OutletGateway {
	db := mysql.Connect(cfg)

	return &OutletGateway{
		OutletRepo:      mysql.NewOutletRepo(db),
		TransactionRepo: mysql.NewDbtransactionRepo(db),
	}
}
