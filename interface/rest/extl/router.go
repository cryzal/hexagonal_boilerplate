package extl

import (
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/protocol/rest"
)

type Controller struct {
	HTTPHandler rest.EchoHTTPHandler
	Config      *config.Config
}

func (r *Controller) RegisterRouter() {
}
