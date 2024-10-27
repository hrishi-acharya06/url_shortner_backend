package model

type ShortUrlRequest struct {
	Url string `json:"url" binding:"required"`
}
