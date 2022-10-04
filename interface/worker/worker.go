package worker

import (
	PortOutlet "hexagonal_boilerplate/core/port/outlet"
	"hexagonal_boilerplate/interface/rest/intl/outlet"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/messaging"
	"hexagonal_boilerplate/shared/models/event"
)

type Routes struct {
	Config     *config.Config
	PortOutlet PortOutlet.Service
	Messaging  messaging.Subscriber
}

func (r *Routes) RegisterRouter() {
	outletHandler := outlet.New(r.PortOutlet)

	r.Messaging.Handle(event.AllEvent, outletHandler.Update())
	r.Messaging.Handle(event.AllEvent2, outletHandler.Update2())
}
