package vehicle

import (
	"context"

	"github.com/8run0/pubsub/pkg/events"
)

type Listener struct {
	S events.Subscriber
}

func (l *Listener) SubscribeForVehicleEvents() {
	l.S.Subscribe(context.Background(), []string{Channel})
}

func (l *Listener) OnArrival(f func(ae ArrivalEvent)) {
	l.S.Register(Channel, Event{
		Payload: &ArrivalEvent{},
	}, func(ctx context.Context, e events.Event) error {
		payload := e.GetPayload()
		arrivalEvent, _ := payload.(*ArrivalEvent)
		f(*arrivalEvent)
		return nil
	})
}
func (l *Listener) OnDeparture(f func(DepartureEvent)) {
	l.S.Register(Channel, Event{
		Payload: &DepartureEvent{},
	}, func(ctx context.Context, e events.Event) error {
		payload := e.GetPayload()
		departureEvent, _ := payload.(*DepartureEvent)
		f(*departureEvent)
		return nil
	})
}
