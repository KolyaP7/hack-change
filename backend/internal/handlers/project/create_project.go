package project

import (
	"encoding/json"
	"hack-change-backend/internal/middleware"
	"hack-change-backend/internal/repository/db"
	"log"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateProject called")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		name string
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "error: unable to get user id from context", http.StatusUnauthorized)
		return
	}

	err := db.CreateProject(req.name, userId)
	if err != nil {
		http.Error(w, "failed to create project", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
