package model

import "gorm.io/gorm"

type UrlMapping struct {
	gorm.Model
	EncodedUrl string
	OrignalUrl string
}

func (UrlMapping) TableName() string {
	return "tabUrl Mapping"
}
