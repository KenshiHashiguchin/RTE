package kvs

import (
	"context"
	"github.com/RTE/web/api/util"
	"github.com/go-redis/redis/v8"
)

var conn *redis.Client
var ctx = context.Background()

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:     util.Env("REDIS_HOST") + ":" + util.Env("REDIS_PORT"),
		Password: util.Env("REDIS_PASSWORD"),
	})
}
