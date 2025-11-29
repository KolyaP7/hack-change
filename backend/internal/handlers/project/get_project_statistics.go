package project

import (
	"encoding/json"
	"log"
	"net/http"

	"hack-change-backend/internal/repository/db"
)

func GetProjectStatistics(w http.ResponseWriter, r *http.Request) {
	log.Println("GetProjectStatistics called")
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}
	projectName := r.URL.Query().Get("project_name")
	if projectName == "" {
		http.Error(w, "project_name is required", http.StatusBadRequest)
		return
	}

	projectStatistics, err := db.GetProjectStatisticsByProjectName(projectName)
	if err != nil {
		http.Error(w, "failed to get project statistics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projectStatistics)
}
