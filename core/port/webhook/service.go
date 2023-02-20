package webhook

type ShopeeWebhookReq struct {
	Data      ShopeeData
	ShopID    int
	Code      int
	Timestamp int
}
type ShopeeData struct {
	Items      []interface{}
	Ordersn    string
	Status     string
	UpdateTime int
}

type Service interface {
	WebhookOrder(provider string, payload interface{}) error
}
