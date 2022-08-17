package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Msg struct {
	Topic string
	Body  []byte
}

type Manager interface {
	SendMsg(ctx context.Context, msg *Msg) error
}

type StreamMQ struct {
	client *redis.Client

	name   string
	maxLen int64
	approx bool
}

type StreamMQOption struct {
	name string
}

type StreamMQOptionApplier func(*StreamMQOption)

func WithStreamName(name string) StreamMQOptionApplier {
	return func(s *StreamMQOption) {
		s.name = name
	}
}

func NewStreamMQ() *StreamMQ {
	return &StreamMQ{}
}

func (q *StreamMQ) SendMsg(ctx context.Context, msg *Msg) error {
	return q.client.XAdd(ctx, &redis.XAddArgs{
		Stream: msg.Topic,
		MaxLen: q.maxLen,
		Approx: q.approx,
		ID:     "*",
		Values: []interface{}{"body", msg.Body},
	}).Err()
}
