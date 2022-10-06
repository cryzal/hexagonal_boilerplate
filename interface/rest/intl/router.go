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

	outletRoute := r.HTTPHandler.Framework.Group("/outlets")
	outletRoute.Use(r.Auth())
	outletRoute.GET("/:outlet_id", outletHandler.Get)
	outletRoute.POST("", outletHandler.Create)
	outletRoute.PUT("/:outlet_id", outletHandler.Edit)
	outletRoute.PUT("/publisher/:outlet_id", outletHandler.EditPublisher)

}
