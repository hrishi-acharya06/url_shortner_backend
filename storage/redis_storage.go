package storage

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func (rc RedisConfig) getDsn() string {
	return fmt.Sprintf("%s:%s", rc.Host, rc.Port)
}
func GetRedisClient() *redis.Client {
	ctx := context.Background()
	redisDb, _ := strconv.Atoi(os.Getenv("redis_db"))
	redisConfig := RedisConfig{
		Host:     os.Getenv("redis_host"),
		Port:     os.Getenv("redis_port"),
		Password: os.Getenv("redis_password"),
		Db:       redisDb,
	}
	addr := redisConfig.getDsn()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.Password,
		DB:       0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("Redis connection successful %s", pong)
	redisClient.Set(ctx, "name", "hrishi", time.Hour)
	return redisClient
}
