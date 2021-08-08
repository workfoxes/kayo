package worker

import "github.com/go-redis/redis"

const (
	KayoDefaultChannel = ""
)

type Worker struct {
	Name      string
	Subscribe *redis.PubSub
}
