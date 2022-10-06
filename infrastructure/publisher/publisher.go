package publisher

import "hexagonal_boilerplate/shared/messaging"

type MessagePublisher struct {
	Publisher messaging.Publisher
}

func (m *MessagePublisher) PublishMessage(topic string, obj any) error {
	err := m.Publisher.Publish(topic, messaging.Payload{Data: obj})

	return err
}
