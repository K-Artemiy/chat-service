package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"chat-service/internal/domain"
)

var (
	ErrChatNotFound    = errors.New("чат не найден")
	ErrInvalidTitle    = errors.New("некорректный заголовок")
	ErrInvalidText     = errors.New("некорректный текст")
	ErrLimitOutOfRange = errors.New("некорректное значение limit")
)

type ChatUsecase struct {
	chats    ChatRepository
	messages MessageRepository
	now      func() time.Time
}

func NewChatUsecase(
	chats ChatRepository,
	messages MessageRepository,
	now func() time.Time,
) *ChatUsecase {
	return &ChatUsecase{chats: chats, messages: messages, now: now}
}

func (uc *ChatUsecase) CreateChat(ctx context.Context, title string) (*domain.Chat, error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 || len(title) > 200 {
		return nil, ErrInvalidTitle
	}
	return uc.chats.Create(ctx, title, uc.now())
}

func (uc *ChatUsecase) SendMessage(ctx context.Context, chatID int64, text string) (*domain.Message, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 || len(text) > 5000 {
		return nil, ErrInvalidText
	}

	chat, err := uc.chats.GetByID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	if chat == nil {
		return nil, ErrChatNotFound
	}

	return uc.messages.Create(ctx, chatID, text, uc.now())
}

func (uc *ChatUsecase) GetChatWithMessages(ctx context.Context, chatID int64, limit int) (*domain.Chat, []domain.Message, error) {
	if limit <= 0 || limit > 100 {
		return nil, nil, ErrLimitOutOfRange
	}

	chat, err := uc.chats.GetByID(ctx, chatID)
	if err != nil {
		return nil, nil, err
	}
	if chat == nil {
		return nil, nil, ErrChatNotFound
	}

	msgs, err := uc.messages.GetLastByChatID(ctx, chatID, limit)
	if err != nil {
		return nil, nil, err
	}
	return chat, msgs, nil
}

func (uc *ChatUsecase) DeleteChat(ctx context.Context, chatID int64) error {
	return uc.chats.DeleteByID(ctx, chatID)
}
