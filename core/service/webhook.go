package service

import (
	"fmt"
	port "hexagonal_boilerplate/core/port/webhook"
	"hexagonal_boilerplate/shared/utils/conv"
)

type Webhook struct {
}

func NewWebhookService() *Webhook {
	return &Webhook{}
}

func (w Webhook) WebhookOrder(provider string, payload interface{}) error {
	switch provider {
	case "SHOPEE":
		var payloadWebhook port.ShopeeWebhookReq
		conv.PassToStruct(payload, &payloadWebhook)

		fmt.Println(payloadWebhook)
	}
	//TODO implement me
	panic("implement me")
	return nil
}
