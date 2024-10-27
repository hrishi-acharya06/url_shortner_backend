package controller

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"
	storage "url_shortner/storage"
	utils "url_shortner/utils"
)

func GenerateShortUrl(orignalUrl string) (int, string) {
	ctx := context.Background()
	urlHashBytes := utils.Sha256Of(orignalUrl)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	encodedStr := utils.Base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))[:8]
	host := os.Getenv("site_host") + "v1/"
	shortUrl := host + encodedStr
	storage.RedisClient.Set(ctx, encodedStr, orignalUrl, time.Hour*1)
	return 200, shortUrl
}
