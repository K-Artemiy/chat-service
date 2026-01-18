package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"chat-service/internal/domain"
	"chat-service/internal/usecase"
)

type ChatGormRepository struct {
	db *gorm.DB
}

func NewChatGormRepository(db *gorm.DB) *ChatGormRepository {
	return &ChatGormRepository{db: db}
}

func (r *ChatGormRepository) Create(ctx context.Context, title string, now time.Time) (*domain.Chat, error) {
	m := ChatModel{Title: title, CreatedAt: now}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return nil, err
	}
	return &domain.Chat{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (r *ChatGormRepository) GetByID(ctx context.Context, id int64) (*domain.Chat, error) {
	var m ChatModel
	err := r.db.WithContext(ctx).First(&m, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usecase.ErrChatNotFound
	}
	if err != nil {
		return nil, err
	}
	return &domain.Chat{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (r *ChatGormRepository) DeleteByID(ctx context.Context, id int64) error {
	res := r.db.WithContext(ctx).Delete(&ChatModel{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return usecase.ErrChatNotFound
	}
	return nil
}

type MessageGormRepository struct {
	db *gorm.DB
}

func NewMessageGormRepository(db *gorm.DB) *MessageGormRepository {
	return &MessageGormRepository{db: db}
}

func (r *MessageGormRepository) Create(ctx context.Context, chatID int64, text string, now time.Time) (*domain.Message, error) {
	m := MessageModel{ChatID: chatID, Text: text, CreatedAt: now}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return nil, err
	}
	return &domain.Message{
		ID:        m.ID,
		ChatID:    m.ChatID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (r *MessageGormRepository) GetLastByChatID(ctx context.Context, chatID int64, limit int) ([]domain.Message, error) {
	var ms []MessageModel
	err := r.db.WithContext(ctx).
		Where("chat_id = ?", chatID).
		Order("created_at DESC").
		Limit(limit).
		Find(&ms).Error
	if err != nil {
		return nil, err
	}
	res := make([]domain.Message, len(ms))

	for i, m := range ms {
		res[i] = domain.Message{
			ID:        m.ID,
			ChatID:    m.ChatID,
			Text:      m.Text,
			CreatedAt: m.CreatedAt,
		}
	}
	return res, nil
}
