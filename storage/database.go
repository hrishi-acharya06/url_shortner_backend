package storage

import (
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var RedisClient *redis.Client
var MysqlClinet *gorm.DB

func InitializeStorage() {
	RedisClient = GetRedisClient()
	MysqlClinet = GetMysqlClient()
}
