package storage

import "url_shortner/model"

func MigrateDatabase() {
	var urlMapping model.UrlMapping
	MysqlClinet.AutoMigrate(&urlMapping)
}
