package webhook

type ShopeeReq struct {
	Data      ShopeeData `json:"data"`
	ShopID    int        `json:"shop_id"`
	Code      int        `json:"code"`
	Timestamp int        `json:"timestamp"`
}
type ShopeeData struct {
	Items      []interface{} `json:"items"`
	Ordersn    string        `json:"ordersn"`
	Status     string        `json:"status"`
	UpdateTime int           `json:"update_time"`
}
