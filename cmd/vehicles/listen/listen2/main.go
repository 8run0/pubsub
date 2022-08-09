package main

import (
	"fmt"
	"time"

	"github.com/8run0/pubsub/pkg/events"
	"github.com/8run0/pubsub/pkg/events/vehicle"
	"github.com/go-redis/redis/v9"
)

func main() {
	listenerClient2 := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,
		ReadTimeout: -1, // use default DB
	})
	listener2 := &vehicle.Listener{
		S: &vehicle.Subscriber{
			Client:     listenerClient2,
			HandlerMap: &events.ChannelEventHandlerMap{},
		},
	}
	listener2.SubscribeForVehicleEvents()
	listener2.OnArrival(func(ae vehicle.ArrivalEvent) {
		fmt.Printf("Arrival Time:\t%s|Registration:\t%s\t--2\n", ae.GetTime().Format(time.RFC850), ae.GetVehicle().Registration)
	})

	listener2.OnDeparture(func(ae vehicle.DepartureEvent) {
		fmt.Printf("Departure Time:\t%s|Registration:\t%s\t--2\n", ae.GetTime().Format(time.RFC850), ae.GetVehicle().Registration)
	})

	time.Sleep(time.Hour)

}
