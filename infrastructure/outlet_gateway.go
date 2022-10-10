package infrastructure

import (
	"hexagonal_boilerplate/infrastructure/publisher"
	"hexagonal_boilerplate/infrastructure/repository/mongo"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/messaging"
)

type OutletGateway struct {
	//OutletRepo      *mysql.RepositoryOutlet
	OutletRepo *mongo.RepositoryOutlet
	//TransactionRepo *mysql.RepositoryDbtransaction
	TransactionRepo *mongo.RepositoryDbtransaction
	Publisher       *publisher.MessagePublisher
}

func NewOutletGateway(cfg *config.Config) *OutletGateway {

	/// example mysql
	//db := mysql.Connect(cfg)

	/// example mongo
	db := mongo.Connect(cfg)

	/// rabbitmq example
	pub := messaging.NewPublisher("amqp://" + cfg.Rabbitmq.User + ":" + cfg.Rabbitmq.Pass + "@" + cfg.Rabbitmq.Host + "")

	return &OutletGateway{
		//OutletRepo:      mysql.NewOutletRepo(db),
		OutletRepo:      mongo.NewOutletRepo(db, cfg.Database.Mongodb.DBName),
		TransactionRepo: mongo.NewDbtransactionRepo(db),
		Publisher:       &publisher.MessagePublisher{pub},
	}
}
