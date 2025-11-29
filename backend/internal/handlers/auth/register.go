package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"hack-change-backend/internal/repository/db"
	"hack-change-backend/pkg/models"
)

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register called")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "name, email, password are required", http.StatusBadRequest)
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("GenerateFromPassword error:", err)
		http.Error(w, "failed to genereta password hash", http.StatusInternalServerError)
		return
	}

	err = db.CreateUser(&models.User{
		UserName:     req.Name,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		log.Println("CreateUser error:", err)
		http.Error(w, fmt.Sprintf("failed to create user: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
