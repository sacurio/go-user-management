package http

import (
	"encoding/json"
	"net/http"
	"user-management/internal/config"
	"user-management/internal/service"

	"github.com/sirupsen/logrus"
)

type UserHTTPHandler struct {
	service *service.UserService
}

func NewUserHTTPHandler(userService *service.UserService) *UserHTTPHandler {
	config.Log.Info("UserHTTPHandler initialized")
	return &UserHTTPHandler{service: userService}
}

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		config.Log.WithField("error", err.Error()).Error("Failed to parse request body")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(request.Name, request.Email)
	if err != nil {
		config.Log.WithField("error", err.Error()).Error("Failed to create user")
		http.Error(w, "Failed to create a new user", http.StatusInternalServerError)
		return
	}

	config.Log.WithFields(logrus.Fields{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}).Info("User created successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHTTPHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		config.Log.WithField("error", err.Error()).Error("Failed to retrieve user list")
		http.Error(w, "Failed to retrieve user list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
