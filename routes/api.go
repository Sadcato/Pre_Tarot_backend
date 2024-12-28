package routes

import (
	"TAROT/pkg/middleware"
	"TAROT/service/chat"
	"encoding/json"
	"net/http"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func RegisterRoutes(chatService *chat.Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/app/readings", handleChat(chatService))

	handler := middleware.LoggingMiddleware(mux)
	handler = middleware.RecoveryMiddleware(handler)

	return handler
}

func handleChat(chatService *chat.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := chatService.Chat(r.Context(), req.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ChatResponse{
			Response: response,
		})
	}
}
