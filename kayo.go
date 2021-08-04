package main

import (
	"context"
	"fmt"
	"github.com/workfoxes/gobase/pkg/client/redis"
	"github.com/workfoxes/gobase/pkg/worker"
)

var ctx = context.Background()

func StartKayo() {
	_ = worker.New("Kayo")
	subscribe := redis.R.Subscribe(ctx, "")
	subscriptions := subscribe.ChannelWithSubscriptions(ctx, 1)
	for {
		select {
		case sub := <-subscriptions:
			fmt.Println(sub)
		}
	}
}

func main() {
	StartKayo()
}
