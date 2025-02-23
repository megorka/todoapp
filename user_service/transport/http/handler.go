package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/megorka/todoapp/user_service/models"
	"github.com/megorka/todoapp/user_service/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req *models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		http.Error(w, "All fields required", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(req.Username, req.Email, req.Password); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("Received signup request: username=%s, email=%s, password=%s", req.Username, req.Email, req.Password)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User successfuly created"})
}

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		log.Printf("Failed to get user by email: %v", err)
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user.Email == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
