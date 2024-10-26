package storage

import (
	redis "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitializeStorage() {
	RedisClient = GetRedisClient()
}
