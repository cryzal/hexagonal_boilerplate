package intl

import (
	"hexagonal_boilerplate/interface/rest/intl/outlet"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/protocol/rest"

	PortOutlet "hexagonal_boilerplate/core/port/outlet"
)

type Routes struct {
	HTTPHandler rest.EchoHTTPHandler
	Config      *config.Config
	PortOutlet  PortOutlet.Service
}

func (r *Routes) RegisterRouter() {
	outletHandler := outlet.New(r.PortOutlet)

	merchantRoute := r.HTTPHandler.Framework.Group("/outlets")
	merchantRoute.Use(r.Auth())
	merchantRoute.GET("/:outlet_id", outletHandler.Get)
}
