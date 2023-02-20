package extl

import (
	port "hexagonal_boilerplate/core/port/webhook"
	"hexagonal_boilerplate/interface/rest/extl/webhook"
	"hexagonal_boilerplate/shared/config"
	"hexagonal_boilerplate/shared/protocol/rest"
)

type Controller struct {
	HTTPHandler rest.EchoHTTPHandler
	Config      *config.Config
	PortWebhook port.Service
}

func (r *Controller) RegisterRouter() {
	webhookHandler := webhook.Newhandler(r.PortWebhook)

	webhookRoute := r.HTTPHandler.Framework.Group("/webhook")
	webhookRoute.POST("/shopee", webhookHandler.Shopee)
}
