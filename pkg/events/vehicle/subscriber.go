package vehicle

import (
	"context"

	"github.com/8run0/pubsub/pkg/events"
	"github.com/go-redis/redis/v9"
)

var _ events.Subscriber = Subscriber{}

type Subscriber struct {
	Client     *redis.Client
	HandlerMap *events.ChannelEventHandlerMap
}

// Register implements events.Subscriber
func (s Subscriber) Register(channel string, e events.Event, handler events.EventHandlerFunc) {
	s.HandlerMap.PutHandler(channel, e, handler)
}

// Subscribe implements events.Subscriber
func (s Subscriber) Subscribe(ctx context.Context, channels []string) {
	subs := s.Client.Subscribe(ctx, channels...)

	go func() {
		for {
			msg := <-subs.Channel()
			event := &Event{
				Payload: nil,
			}
			event.UnmarshalBinary([]byte(msg.Payload))
			handlers := s.HandlerMap.GetHandlersForChannelAndEvent(msg.Channel, event)
			for _, h := range handlers {
				h(ctx, event)
			}
		}
	}()
}
