package dao

import (
	"sync"

	"github.com/go-redis/redis"
)

const (
	Host     = "47.240.95.242:6379"
	Password = "www.ak123.com"
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
