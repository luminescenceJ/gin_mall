package conf

import (
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
