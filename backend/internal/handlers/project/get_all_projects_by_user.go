package project

import (
	"encoding/json"
	"hack-change-backend/internal/middleware"
	"hack-change-backend/internal/repository/db"
	"log"
	"net/http"
)

func GetProjectsByUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllProjectsByUser called")
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "error: unable to get user id from context", http.StatusUnauthorized)
		return
	}

	projects, err := db.GetProjectsById(userId)
	if err != nil {
		http.Error(w, "failed to get all projects by user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
