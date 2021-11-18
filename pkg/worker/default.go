package worker

import "github.com/go-redis/redis/v8"

const (
	KayoDefaultChannel = ""
)

type Worker struct {
	Name      string
	Subscribe *redis.PubSub
}
