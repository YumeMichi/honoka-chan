package database

import (
	"context"
	"honoka-chan/config"

	"github.com/redis/go-redis/v9"
)

var (
	RedisCli *redis.Client
	RedisCtx = context.Background()
)

func init() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Host + ":" + config.Conf.Redis.Port,
		Password: config.Conf.Redis.Pass,
		DB:       config.Conf.Redis.Db,
	})

	_, err := RedisCli.Ping(RedisCtx).Result()
	if err != nil {
		panic(err)
	}
}
