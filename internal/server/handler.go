package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"chat-service/internal/usecase"
)

type ChatHandler struct {
	uc  *usecase.ChatUsecase
	log *log.Logger
}

func NewChatHandler(uc *usecase.ChatUsecase, logger *log.Logger) *ChatHandler {
	return &ChatHandler{uc: uc, log: logger}
}

func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	chat, err := h.uc.CreateChat(r.Context(), req.Title)
	if err != nil {
		if err == usecase.ErrInvalidTitle {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.internalError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, chat)
}

func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	chatID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	var req messageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	msg, err := h.uc.SendMessage(r.Context(), chatID, req.Text)
	if err != nil {
		switch err {
		case usecase.ErrInvalidText:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case usecase.ErrChatNotFound:
			http.Error(w, "чат не найден", http.StatusNotFound)
		default:
			h.internalError(w, err)
		}
		return
	}

	writeJSON(w, http.StatusCreated, msg)
}

func (h *ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	chatID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	limit := 20
	if l := r.URL.Query().Get("limit"); l != "" {
		v, err := strconv.Atoi(l)
		if err != nil {
			http.Error(w, "некорректный limit", http.StatusBadRequest)
			return
		}
		limit = v
	}

	chat, msgs, err := h.uc.GetChatWithMessages(r.Context(), chatID, limit)
	if err != nil {
		switch err {
		case usecase.ErrChatNotFound:
			http.Error(w, "чат не найден", http.StatusNotFound)
		case usecase.ErrLimitOutOfRange:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			h.internalError(w, err)
		}
		return
	}

	resp := chatResponse{
		Chat:     chat,
		Messages: msgs,
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *ChatHandler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	chatID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "некорректный id", http.StatusBadRequest)
		return
	}

	if err := h.uc.DeleteChat(r.Context(), chatID); err != nil {
		if err == usecase.ErrChatNotFound {
			http.Error(w, "чат не найден", http.StatusNotFound)
			return
		}
		h.internalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChatHandler) internalError(w http.ResponseWriter, err error) {
	h.log.Println("internal error:", err)
	http.Error(w, "внутренняя ошибка сервера", http.StatusInternalServerError)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
