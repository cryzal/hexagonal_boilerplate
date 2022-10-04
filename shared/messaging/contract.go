package messaging

type Publisher interface {
	Publish(topic string, delayInMS int, payload Payload) error
}

type HandleFunc func(payload Payload, err error)

type Subscriber interface {
	Handle(topic string, onReceived HandleFunc)
	Run(url string)
}
