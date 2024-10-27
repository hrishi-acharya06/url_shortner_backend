package controller

import (
	"context"
	storage "url_shortner/storage"
)

func GetOrignalUrl(shortUrl string) string {
	ctx := context.Background()
	orignalUrl, _ := storage.RedisClient.Get(ctx, shortUrl).Result()
	if orignalUrl == "" {
		panic("Orignal Url not found")
	}
	return orignalUrl
}
