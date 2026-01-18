package usecase

import (
	"context"
	"time"

	"chat-service/internal/domain"
)

type ChatRepository interface {
	Create(ctx context.Context, title string, now time.Time) (*domain.Chat, error)
	GetByID(ctx context.Context, id int64) (*domain.Chat, error)
	DeleteByID(ctx context.Context, id int64) error
}

type MessageRepository interface {
	Create(ctx context.Context, chatID int64, text string, now time.Time) (*domain.Message, error)
	GetLastByChatID(ctx context.Context, chatID int64, limit int) ([]domain.Message, error)
}
