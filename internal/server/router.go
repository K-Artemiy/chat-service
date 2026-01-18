package server

import "net/http"

func NewRouter(h *ChatHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /chats", h.CreateChat)
	mux.HandleFunc("POST /chats/{id}/messages", h.SendMessage)
	mux.HandleFunc("GET /chats/{id}", h.GetChat)
	mux.HandleFunc("DELETE /chats/{id}", h.DeleteChat)

	return mux
}
