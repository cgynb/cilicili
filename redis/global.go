package redis

import (
	"cilicili/config"
	"github.com/go-redis/redis"
)

var DB *redis.Client

func Init() {
	DB = redis.NewClient(&redis.Options{
		Addr: config.Conf.RedisConfig.Host + ":" + config.Conf.RedisConfig.Port,
		DB:   config.Conf.RedisConfig.DB,
	})
	_, err := DB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
