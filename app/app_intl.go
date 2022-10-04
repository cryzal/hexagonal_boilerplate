package app

import (
	"hexagonal_boilerplate/infrastructure"
	"hexagonal_boilerplate/interface/rest/intl"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/driver"
	"hexagonal_boilerplate/shared/protocol/rest"

	"hexagonal_boilerplate/core/service"
)

type Intl struct {
	httpHandler *rest.EchoHTTPHandler
	router      driver.Router
}

func (c Intl) RunApplication() {
	c.router.RegisterRouter()
	c.httpHandler.RunApplication()
}
func NewIntl() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig("APP_INTL_ADDRESS")

		httpHandler := rest.NewEchoHTTPHandlerDefault(cfg)
		datasource := infrastructure.NewOutletGateway(cfg)

		return &Intl{
			httpHandler: &httpHandler,
			router: &intl.Routes{
				HTTPHandler: httpHandler,
				Config:      cfg,
				PortOutlet:  service.OutletServiceNew(datasource.OutletRepo),
			},
		}

	}
}
