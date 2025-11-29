package project

import (
	"encoding/json"
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
}
