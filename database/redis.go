package database

import (
	"context"
	"errors"
	"honoka-chan/config"
	"strconv"

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

func GetUid(key string) (int, error) {
	uid, err := RedisCli.HGet(RedisCtx, "login_key_uid", key).Result()
	if err != nil {
		return 0, err
	}
	userId, err := strconv.Atoi(uid)
	if err != nil {
		return 0, err
	}

	if userId == 0 {
		return 0, errors.New("userId is 0")
	}

	return userId, nil
}

func MatchTokenUid(token, uid string) bool {
	res, err := RedisCli.HGet(RedisCtx, "token_uid", token).Result()
	if err != nil {
		panic(err)
	}

	return res == uid
}
