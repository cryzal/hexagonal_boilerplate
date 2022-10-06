package publisher

type CreateOutlet struct {
	OutletID string `json:"outlet_id,omitempty"`
}

type Publisher interface {
	PublishMessage(topic string, obj any) error
}
