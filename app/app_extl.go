package app

import (
	"hexagonal_boilerplate/interface/rest/extl"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/driver"
	"hexagonal_boilerplate/shared/protocol/rest"
)

type Extl struct {
	httpHandler *rest.EchoHTTPHandler
	controller  driver.Router
}

func (c Extl) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}
func NewExtl() func() driver.RegistryContract {
	return func() driver.RegistryContract {
		cfg := config.ReadConfig("APP_EXTL_ADDRESS")

		httpHandler := rest.NewEchoHTTPHandlerDefault(cfg)

		return &Extl{
			httpHandler: &httpHandler,
			controller: &extl.Controller{
				HTTPHandler: httpHandler,
				Config:      cfg,
			},
		}

	}
}
