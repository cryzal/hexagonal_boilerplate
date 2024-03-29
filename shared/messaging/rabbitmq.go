package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var exchangeName = "boilerplate"
var exchangeType = "direct"

type publisherImpl struct {
	rabbitMQChannel *amqp.Channel
}

// NewPublisher is
// url "amqp://guest:guest@localhost:5672/"
func NewPublisher(url string) *publisherImpl {

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil
	}
	//defer ch.Close()

	return &publisherImpl{
		rabbitMQChannel: ch,
	}
}

// Publish is
func (m *publisherImpl) Publish(topic string, data Payload) error {

	dataInBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = m.rabbitMQChannel.PublishWithContext(
		ctx,
		exchangeName, // exchange
		topic,        // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataInBytes,
		})
	if err != nil {
		return err
	}

	return nil
}

type subscriberImpl struct {
	queueName string
	topicMap  map[string]HandleFunc
}

// NewSubscriber is
func NewSubscriber(queueName string) Subscriber {
	return &subscriberImpl{
		queueName: queueName,
		topicMap:  map[string]HandleFunc{},
	}
}

func (r *subscriberImpl) Handle(topic string, onReceived HandleFunc) {

	r.topicMap[topic] = onReceived

}

// Run is
// "amqp://guest:guest@localhost:5672/"
func (r *subscriberImpl) Run(url string) {

	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err.Error())
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	rabbitMQChannel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		err := rabbitMQChannel.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	args := amqp.Table{
		"x-delayed-type": "topic", // only for x-delay-message
	}

	err = rabbitMQChannel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		args,         // arguments
	)
	if err != nil {
		panic(err.Error())
	}

	for s := range r.topicMap {
		q, err := rabbitMQChannel.QueueDeclare(
			r.queueName+"-"+s, // name
			false,             // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // no-wait
			nil,               // arguments
		)
		if err != nil {
			panic(err.Error())
		}

		err = rabbitMQChannel.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			exchangeName, // exchange
			false,
			nil,
		)
		if err != nil {
			panic(err.Error())
		}

		deliveryMsg, err := rabbitMQChannel.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Start worker. Exchange Name : %s. queue name : %s. Routing-key : %s\n", exchangeName, q.Name, s)

		go func(routingKey string) {

			for d := range deliveryMsg {
				var data Payload
				err := json.Unmarshal(d.Body, &data)
				log.Printf("Receive message from %s. Body : %s\n", d.RoutingKey, data)

				r.topicMap[routingKey](data, err)
				//log.Printf("recv %s %s", d.RoutingKey, data.Data)
			}
		}(s)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
}
