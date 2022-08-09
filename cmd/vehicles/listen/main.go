package main

import (
	"fmt"
	"time"

	"github.com/8run0/pubsub/pkg/events"
	"github.com/8run0/pubsub/pkg/events/vehicle"
	"github.com/go-redis/redis/v9"
)

func main() {
	listenerClient := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,
		ReadTimeout: -1, // use default DB
	})
	listener := &vehicle.Listener{
		S: &vehicle.Subscriber{
			Client:     listenerClient,
			HandlerMap: &events.ChannelEventHandlerMap{},
		},
	}
	listener.SubscribeForVehicleEvents()
	listener.OnArrival(func(ae vehicle.ArrivalEvent) {
		fmt.Printf("Arrival Time:\t%s|Registration:\t%s\n", ae.GetTime().Format(time.RFC850), ae.GetVehicle().Registration)
	})

	listener.OnDeparture(func(ae vehicle.DepartureEvent) {
		fmt.Printf("Departure Time:\t%s|Registration:\t%s\n", ae.GetTime().Format(time.RFC850), ae.GetVehicle().Registration)
	})
	time.Sleep(time.Hour)

}
