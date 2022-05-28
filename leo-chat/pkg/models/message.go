package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	Text       string `json:"text"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
}

func (c *Message) CreateMessage() *Message {
	db.Create(&c)
	return c
}
