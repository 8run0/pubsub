package vehicle

import (
	"context"
	"time"
)

type Notifier struct {
	P Publisher
}

func (n Notifier) Arrival(ctx context.Context, v *Vehicle) {
	n.P.Publish(ctx, ArrivalEvent{
		Vehicle: v,
		Time:    time.Now(),
	})
}
func (n Notifier) Departure(ctx context.Context, v *Vehicle) {
	n.P.Publish(ctx, DepartureEvent{
		Vehicle: v,
		Time:    time.Now(),
	})
}
