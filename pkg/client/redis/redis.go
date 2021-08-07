package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/workfoxes/kayo/pkg/errors"
	"sync"
	"time"
)

var r Redis

type Redis struct {
	once     sync.Once
	_default *redis.Client
}

func (r *Redis) Init() {
	var err error = nil
	r.once.Do(func() {
		r._default = redis.NewClient(&redis.Options{
			Addr:     "", //config.RedisHost,
			Password: "", //config.RedisPassword, // no password set
			DB:       0,  // use default DB
		})
		if err = r._default.Ping(nil).Err(); err != nil {
			panic(errors.RedisUnreachable)
		}
	})
}

func Get(ctx context.Context, key string) *redis.StringCmd {
	r.Init()
	return r._default.Get(ctx, key)
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	r.Init()
	return r._default.Set(ctx, key, value, expiration)
}
