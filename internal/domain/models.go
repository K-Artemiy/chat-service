package domain

import "time"

type Chat struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}

type Message struct {
	ID        int64
	ChatID    int64
	Text      string
	CreatedAt time.Time
}
