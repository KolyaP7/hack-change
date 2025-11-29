package files

import (
	"log"
	"net/http"
	"strconv"

	"hack-change-backend/internal/logic"
	"hack-change-backend/internal/middleware"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadFile called")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "multipart/form-data" {
		http.Error(w, "Content-Type must be multipart/form-data", http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "error: unable to get user id from context", http.StatusUnauthorized)
		return
	}

	projectId, err := strconv.Atoi(r.FormValue("project_id"))

	if err != nil {
		http.Error(w, "project_id must be int", http.StatusBadRequest)
		return
	}

	file, ok := r.MultipartForm.File["file"]
	if !ok {
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}

	logic.FileToDB(file[0], userId, projectId)

}
