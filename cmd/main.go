package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"chat-service/internal/config"
	chatrepo "chat-service/internal/repository"
	httphandler "chat-service/internal/server"
	"chat-service/internal/usecase"
)

func main() {
	logger := log.New(os.Stdout, "[chat-api] ", log.LstdFlags|log.Lshortfile)

	cfg := config.Load()
	if cfg.DSN == "" {
		logger.Fatal("DB_DSN не задан")
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		logger.Fatalf("ошибка подключения к БД: %v", err)
	}

	chatRepo := chatrepo.NewChatGormRepository(db)
	msgRepo := chatrepo.NewMessageGormRepository(db)

	uc := usecase.NewChatUsecase(chatRepo, msgRepo, func() time.Time {
		return time.Now().UTC()
	})

	handler := httphandler.NewChatHandler(uc, logger)
	router := httphandler.NewRouter(handler)

	logger.Printf("запуск HTTP-сервера на %s\n", cfg.Addr)

	if err := http.ListenAndServe(cfg.Addr, router); err != nil {
		logger.Fatalf("ошибка сервера: %v", err)
	}
}
