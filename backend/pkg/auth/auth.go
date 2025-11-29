package auth

import (
	"fmt"
	"hack-change-backend/pkg/getenv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(getenv.GetValue("JWT_SECRET", "JWT"))

// tokenBlacklist хранит зачеркнутые токены
// Используем map для быстрого поиска и sync.RWMutex для потокобезопасности
var (
	tokenBlacklist = make(map[string]time.Time)
	blacklistMutex sync.RWMutex
)

func VerifyToken(tokenString string) (int, error) {
	// Проверяем, не находится ли токен в черном списке
	blacklistMutex.RLock()
	if expiry, exists := tokenBlacklist[tokenString]; exists {
		blacklistMutex.RUnlock()
		// Если токен в черном списке и еще не истек срок его действия, возвращаем ошибку
		if time.Now().Before(expiry) {
			return 0, fmt.Errorf("token has been revoked")
		}
		// Если срок действия истек, удаляем из черного списка
		blacklistMutex.Lock()
		delete(tokenBlacklist, tokenString)
		blacklistMutex.Unlock()
	} else {
		blacklistMutex.RUnlock()
	}

	// 1. Создаем пустой экземпляр jwt.MapClaims для заполнения
	claims := jwt.MapClaims{}

	// 2. Используем jwt.ParseWithClaims и передаем наш пустой claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		// Здесь уже будут обработаны ошибки, такие как ErrTokenExpired,
		// если токен не прошел проверку по сроку годности (exp)
		return 0, err
	}

	// Проверка на валидность токена
	if !token.Valid {
		return 0, jwt.ErrTokenInvalidClaims
	}

	// 3. Теперь claims уже заполнен, и он является jwt.MapClaims,
	//    проверка 'ok' больше не нужна, и мы можем работать с полями.

	// Вам больше не нужна ручная проверка времени, т.к. ParseWithClaims
	// делает это автоматически и возвращает ErrTokenExpired!

	// Преобразуем userID в int (оставляем вашу логику преобразования)
	var userID int
	switch v := claims["userID"].(type) {
	case float64:
		userID = int(v)
	case string:
		// Если это строка, попробуем преобразовать
		// В реальном коде лучше использовать strconv.Atoi
		if v == "1" {
			userID = 1
		} else {
			return 0, jwt.ErrTokenInvalidClaims
		}
	default:
		return 0, jwt.ErrTokenInvalidClaims
	}
	return userID, nil
}

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

// RevokeToken добавляет токен в черный список
// Токен будет считаться недействительным до истечения срока его действия
func RevokeToken(tokenString string) error {
	// Парсим токен, чтобы получить время истечения
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// Получаем время истечения из claims
	var expiry time.Time
	if exp, ok := claims["exp"].(float64); ok {
		expiry = time.Unix(int64(exp), 0)
	} else {
		// Если не удалось получить время истечения, используем время по умолчанию (1 час)
		expiry = time.Now().Add(1 * time.Hour)
	}

	// Добавляем токен в черный список
	blacklistMutex.Lock()
	tokenBlacklist[tokenString] = expiry
	blacklistMutex.Unlock()

	return nil
}

// cleanupExpiredTokens периодически очищает истекшие токены из черного списка
// Можно вызвать в горутине для автоматической очистки
func CleanupExpiredTokens() {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()

	now := time.Now()
	for token, expiry := range tokenBlacklist {
		if now.After(expiry) {
			delete(tokenBlacklist, token)
		}
	}
}
