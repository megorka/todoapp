package router

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/megorka/todoapp/authorization/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SendUserCreatedEvent(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := h.service.RegisterUser(ctx, req.Username, req.Email, req.Password); err != nil {
		http.Error(w, "Failed to send event", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User event send successfully"})
}