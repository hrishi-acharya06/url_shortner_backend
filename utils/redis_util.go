package utils

import (
	"context"
	"fmt"
	"sync"
	"time"
	storage "url_shortner/storage"
)

func AddToRedis(wg *sync.WaitGroup, key string, value string) {
	defer wg.Done()
	ctx := context.Background()
	err := storage.RedisClient.Set(ctx, key, value, time.Hour*1).Err()
	if err != nil {
		//some logic to log this issue
	}
}

func UpdateKeyCounter(wg *sync.WaitGroup, key string) {
	defer wg.Done()
	ctx := context.Background()
	counterKey := fmt.Sprintf("counter:%s", key)
	err := storage.RedisClient.Incr(ctx, counterKey).Err()
	if err != nil {
		//some logic to log this issue
	}

}
