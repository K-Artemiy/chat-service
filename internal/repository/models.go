package repository

import "time"

type ChatModel struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	Title     string    `gorm:"column:title;type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()"`
}

func (ChatModel) TableName() string { return "chats" }

type MessageModel struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	ChatID    int64     `gorm:"column:chat_id;not null;index"`
	Text      string    `gorm:"column:text;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()"`
}

func (MessageModel) TableName() string { return "messages" }
