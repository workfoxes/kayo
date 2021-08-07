package main

import (
	"context"
)

var ctx = context.Background()

func StartKayo() {
	//_ = worker.New("Kayo")
	//subscribe := redis.R.Subscribe(ctx, "")
	//subscriptions := subscribe.ChannelWithSubscriptions(ctx, 1)
	//for {
	//	select {
	//	case sub := <-subscriptions:
	//		fmt.Println(sub)
	//	}
	//}
}

func main() {
	StartKayo()
}
