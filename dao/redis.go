package dao

import (
	"sync"

	"github.com/go-redis/redis"
)

const (
	Host     = "redis"
	Password = ""
	Db       = 0
)

var redisInstance *redis.Client
var once sync.Once

func RedisClient() *redis.Client {
	once.Do(func() {
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     Host,
			Password: Password,
			DB:       Db,
		})
	})
	return redisInstance
}
