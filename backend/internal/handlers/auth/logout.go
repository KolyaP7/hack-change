package auth

import (
	"log"
	"net/http"
	"strings"

	"hack-change-backend/pkg/auth"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем токен из заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: No or invalid Bearer header", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Добавляем токен в черный список
	if err := auth.RevokeToken(token); err != nil {
		log.Printf("Failed to revoke token: %v\n", err)
		http.Error(w, "Failed to revoke token", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully logged out"}`))
}
