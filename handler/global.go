package handler

import (
	"context"
	"honoka-chan/config"

	"github.com/redis/go-redis/v9"
)

var (
	nonce = 0

	redisCli *redis.Client
	redisCtx = context.Background()
)

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Host + ":" + config.Conf.Redis.Port,
		Password: config.Conf.Redis.Pass,
		DB:       config.Conf.Redis.Db,
	})

	_, err := redisCli.Ping(redisCtx).Result()
	if err != nil {
		panic(err)
	}
}
