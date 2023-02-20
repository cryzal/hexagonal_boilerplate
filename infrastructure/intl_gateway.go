package infrastructure

import (
	"hexagonal_boilerplate/infrastructure/publisher"
	"hexagonal_boilerplate/infrastructure/repository/mysql"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/messaging"
)

type IntlGateway struct {
	OutletRepo *mysql.RepositoryOutlet
	//OutletRepo *mongo.RepositoryOutlet
	TransactionRepo *mysql.RepositoryDbtransaction
	//TransactionRepo *mongo.RepositoryDbtransaction
	Publisher *publisher.MessagePublisher
}

func NewIntlGateway(cfg *config.Config) *IntlGateway {

	/// example mysql
	db := mysql.Connect(cfg)

	/// example mongo
	//db := mongo.Connect(cfg)

	/// rabbitmq example
	pub := messaging.NewPublisher("amqp://" + cfg.Rabbitmq.User + ":" + cfg.Rabbitmq.Pass + "@" + cfg.Rabbitmq.Host + "")

	return &IntlGateway{
		OutletRepo: mysql.NewOutletRepo(db),
		//OutletRepo:      mongo.NewOutletRepo(db, cfg.Database.Mongodb.DBName),
		TransactionRepo: mysql.NewDbtransactionRepo(db),
		Publisher:       &publisher.MessagePublisher{pub},
	}
}
