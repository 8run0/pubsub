package vehicle

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type Publisher struct {
	Client *redis.Client
}

func (p *Publisher) Publish(ctx context.Context, e VehicleEvent) error {
	cmd := p.Client.Publish(ctx, Channel, e)
	_, err := cmd.Result()
	return err
}
