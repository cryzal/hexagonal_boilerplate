package messaging

//
//import (
//	"context"
//	"crypto/tls"
//	"encoding/json"
//	"github.com/go-redis/redis/v8"
//)
//
//var (
//	ctx = context.Background()
//)
//
//type publisherRedisImpl struct {
//	redis *redis.Client
//}
//
//func NewPublisherRedis(addr, username, pass string, tlsMode bool) *publisherRedisImpl {
//	opt := &redis.Options{
//		Addr:      addr,
//		Username:  username,
//		Password:  pass,
//		TLSConfig: nil,
//		DB:        0,
//	}
//
//	if tlsMode {
//		opt.TLSConfig = &tls.Config{
//			InsecureSkipVerify: true,
//		}
//	}
//	redis := redis.NewClient(opt)
//	_, err := redis.Ping(ctx).Result()
//	if err != nil {
//		panic(err)
//	}
//	return &publisherRedisImpl{
//		redis: redis,
//	}
//}
//func (m *publisherRedisImpl) Publish(channel string, data Payload) error {
//	err := m.redis.Publish(ctx, channel, data)
//	if err.Err() != nil {
//		return err.Err()
//	}
//	return nil
//}
//
//type subscriberRedisImpl struct {
//	queueName  string
//	channelMap map[string]HandleFunc
//}
//
//// NewSubscriber is
//func NewSubscriberRedis(channel string) *subscriberRedisImpl {
//	return &subscriberRedisImpl{
//		queueName:  channel,
//		channelMap: map[string]HandleFunc{},
//	}
//}
//func (r *subscriberRedisImpl) Handle(channel string, onReceived HandleFunc) {
//	r.channelMap[channel] = onReceived
//}
//
//func (r *subscriberRedisImpl) Run(addr, username, pass string, tlsMode bool) {
//	opt := &redis.Options{
//		Addr:      addr,
//		Username:  username,
//		Password:  pass,
//		TLSConfig: nil,
//		DB:        0,
//	}
//
//	if tlsMode {
//		opt.TLSConfig = &tls.Config{
//			InsecureSkipVerify: true,
//		}
//	}
//	redis := redis.NewClient(opt)
//	_, err := redis.Ping(ctx).Result()
//	if err != nil {
//		panic(err)
//	}
//	for s := range r.channelMap {
//		subscriber := redis.Subscribe(ctx, s)
//		for {
//			msg, err := subscriber.ReceiveMessage(ctx)
//			if err != nil {
//				panic(err.Error())
//			}
//			var data Payload
//			err = json.Unmarshal([]byte(msg.Payload), &data)
//			r.channelMap[s](data, err)
//		}
//	}
//}
