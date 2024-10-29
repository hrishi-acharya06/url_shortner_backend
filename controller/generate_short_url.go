package controller

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	utils "url_shortner/utils"
)

func GenerateShortUrl(orignalUrl string) (int, string) {
	wg := &sync.WaitGroup{}
	urlHashBytes := utils.Sha256Of(orignalUrl)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	encodedStr := utils.Base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))[:8]
	host := os.Getenv("site_host") + "v1/"
	shortUrl := host + encodedStr
	wg.Add(3)
	go utils.AddToRedis(wg, encodedStr, orignalUrl)
	go utils.UpdateKeyCounter(wg, encodedStr)
	wg.Wait()
	return 200, shortUrl
}
