package events

import (
	"context"
	"reflect"
	"sync"
)

type Event interface {
	GetPayload() any
}

type BaseEvent struct {
	Payload any
}

type Publisher interface {
	Publish(ctx context.Context, e Event) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, channels []string)
	Register(channel string, e Event, handler EventHandlerFunc)
}

// type Handler struct {
// 	HandleFunc EventHandlerFunc
// }

// func (h *Handler) Handle(ctx context.Context, e Event) error {
// 	return h.HandleFunc(ctx, e)
// }

type EventHandlerFunc func(context.Context, Event) error

type ChannelEventHandlerMap struct {
	sync.Map
}

type eventHandlerMap map[string][]EventHandlerFunc

func (e *ChannelEventHandlerMap) GetHandlersForChannelAndEvent(channel string, event Event) []EventHandlerFunc {
	eventHandlerMap := e.getEventHandlerMap(channel)
	eventType := reflect.TypeOf(event.GetPayload()).String()
	if h, ok := eventHandlerMap[eventType]; !ok {
		//no handlers for event
		return []EventHandlerFunc{}
	} else {
		return h
	}
}

func (e *ChannelEventHandlerMap) PutHandler(channel string, event Event, handler EventHandlerFunc) {
	eventHandlerMap := e.getEventHandlerMap(channel)
	// add handler to list of handlers for event type
	var handlers []EventHandlerFunc
	eventType := reflect.TypeOf(event.GetPayload()).String()
	if h, ok := eventHandlerMap[eventType]; !ok {
		//no handlers for event
		handlers = []EventHandlerFunc{}
	} else {
		handlers = h
	}
	handlers = append(handlers, handler)
	eventHandlerMap[eventType] = handlers
	e.Store(channel, eventHandlerMap)
}
func (e ChannelEventHandlerMap) getEventHandlerMap(channel string) (ehmap eventHandlerMap) {
	handlers, ok := e.Load(channel)
	if ok {
		ehmap = handlers.(eventHandlerMap)
	} else {
		ehmap = make(eventHandlerMap, 0)
	}
	return ehmap
}
