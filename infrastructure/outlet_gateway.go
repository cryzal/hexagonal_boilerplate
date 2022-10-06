package infrastructure

import (
	"hexagonal_boilerplate/infrastructure/publisher"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/messaging"

	"hexagonal_boilerplate/infrastructure/repository/mysql"
)

type OutletGateway struct {
	OutletRepo      *mysql.RepositoryOutlet
	TransactionRepo *mysql.RepositoryDbtransaction
	Publisher       *publisher.MessagePublisher
}

func NewOutletGateway(cfg *config.Config) *OutletGateway {
	db := mysql.Connect(cfg)

	/// rabbitmq example
	pub := messaging.NewPublisher("amqp://" + cfg.Rabbitmq.User + ":" + cfg.Rabbitmq.Pass + "@" + cfg.Rabbitmq.Host + "")

	return &OutletGateway{
		OutletRepo:      mysql.NewOutletRepo(db),
		TransactionRepo: mysql.NewDbtransactionRepo(db),
		Publisher:       &publisher.MessagePublisher{pub},
	}
}
