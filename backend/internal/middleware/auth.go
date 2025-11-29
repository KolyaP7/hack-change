package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"hack-change-backend/pkg/auth"

	"github.com/golang-jwt/jwt/v5"
)

const UserIDKey string = "userID"

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized: No or invalid Bearer header", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			userID, err := auth.VerifyToken(token)

			if err != nil {
				var status int
				var errMsg string

				if errors.Is(err, jwt.ErrTokenExpired) {
					status = http.StatusUnauthorized
					errMsg = "Token expired"
				} else if errors.Is(err, jwt.ErrTokenInvalidClaims) || errors.Is(err, jwt.ErrSignatureInvalid) {
					status = http.StatusUnauthorized
					errMsg = "Token invalid"
				} else {
					status = http.StatusInternalServerError
					errMsg = "Internal Server Error"
					log.Printf("Verification error: %v\n", err)
				}

				http.Error(w, errMsg, status)
				return
			}

			// Add logging for successful authentication
			log.Printf("Authentication successful for userID: %d\n", userID)

			// Token is valid. Save userID in context
			// We use the constant key UserIDKey for type safety
			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			// Pass control to the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
