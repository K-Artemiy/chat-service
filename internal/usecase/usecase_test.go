package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"chat-service/internal/domain"
	"chat-service/internal/usecase"
)

type ChatRepoMock struct {
	mock.Mock
}

func (m *ChatRepoMock) Create(ctx context.Context, title string, now time.Time) (*domain.Chat, error) {
	args := m.Called(ctx, title, now)
	if c := args.Get(0); c != nil {
		return c.(*domain.Chat), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ChatRepoMock) GetByID(ctx context.Context, id int64) (*domain.Chat, error) {
	args := m.Called(ctx, id)
	if c := args.Get(0); c != nil {
		return c.(*domain.Chat), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *ChatRepoMock) DeleteByID(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MessageRepoMock struct {
	mock.Mock
}

func (m *MessageRepoMock) Create(ctx context.Context, chatID int64, text string, now time.Time) (*domain.Message, error) {
	args := m.Called(ctx, chatID, text, now)
	if msg := args.Get(0); msg != nil {
		return msg.(*domain.Message), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MessageRepoMock) GetLastByChatID(ctx context.Context, chatID int64, limit int) ([]domain.Message, error) {
	args := m.Called(ctx, chatID, limit)
	if v := args.Get(0); v != nil {
		return v.([]domain.Message), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateChat_TrimAndValidate(t *testing.T) {
	chatRepo := new(ChatRepoMock)
	msgRepo := new(MessageRepoMock)

	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	uc := usecase.NewChatUsecase(chatRepo, msgRepo, func() time.Time { return now })

	chatRepo.
		On("Create", mock.Anything, "Название", now).
		Return(&domain.Chat{ID: 1, Title: "Название", CreatedAt: now}, nil)

	chat, err := uc.CreateChat(context.Background(), "  Название  ")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), chat.ID)
	assert.Equal(t, "Название", chat.Title)
}

func TestSendMessage_InvalidText(t *testing.T) {
	uc := usecase.NewChatUsecase(new(ChatRepoMock), new(MessageRepoMock), time.Now)
	_, err := uc.SendMessage(context.Background(), 1, " ")
	assert.Equal(t, usecase.ErrInvalidText, err)
}
