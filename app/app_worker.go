package app

import (
	"hexagonal_boilerplate/core/service"
	"hexagonal_boilerplate/infrastructure"
	"hexagonal_boilerplate/interface/worker"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/driver"
	"hexagonal_boilerplate/shared/messaging"
)

type app_worker struct {
	Messaging messaging.Subscriber
	Host      string
	Username  string
	Password  string
	router    driver.Router
}

func (a app_worker) RunApplication() {
	//TODO implement me
	a.router.RegisterRouter()
	a.Messaging.Run("amqp://" + a.Username + ":" + a.Password + "@" + a.Host + "")
}

func NewWorker() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig("APP_INTL_ADDRESS")
		datasource := infrastructure.NewOutletGateway(cfg)
		subs := messaging.NewSubscriber("outlet")
		return &app_worker{
			Messaging: subs,
			Host:      cfg.Rabbitmq.Host,
			Username:  cfg.Rabbitmq.User,
			Password:  cfg.Rabbitmq.Pass,
			router: &worker.Routes{
				Config:     cfg,
				PortOutlet: service.OutletServiceNew(datasource.OutletRepo, datasource.TransactionRepo, datasource.Publisher),
				Messaging:  subs,
			},
		}

	}
}
