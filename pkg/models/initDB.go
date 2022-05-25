package models

import (
	"gorm.io/gorm"
	"lavless.com/chat/pkg/config"
)

var db *gorm.DB

func init() {
	config.ConnectDb()
	db = config.GetDB()
	db.AutoMigrate(&Message{})
}