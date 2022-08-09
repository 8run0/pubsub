package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/8run0/pubsub/pkg/events/vehicle"
	"github.com/go-redis/redis/v9"
)

var carregs = []string{"tgz6368", "cfd1123", "abc1234", "cdz3480", "tqa1299", "aqd123"}

func randomReg() string {
	i := rand.Intn(len(carregs))
	return carregs[i]
}

func newArrivals(n *vehicle.Notifier) {
	for {
		ctx := context.Background()
		reg := randomReg()
		n.Arrival(ctx, &vehicle.Vehicle{
			Registration: vehicle.Registration(reg),
		})
		time.Sleep(time.Second * 3)
	}
}
func newDepartures(n *vehicle.Notifier) {
	for {
		ctx := context.Background()
		reg := randomReg()
		n.Departure(ctx, &vehicle.Vehicle{
			Registration: vehicle.Registration(reg),
		})
		time.Sleep(time.Second * 7)
	}
}

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	n := &vehicle.Notifier{
		P: vehicle.Publisher{
			Client: redisClient,
		},
	}
	go func() {
		newArrivals(n)
	}()
	go func() {
		newDepartures(n)
	}()
	time.Sleep(time.Hour)
}
