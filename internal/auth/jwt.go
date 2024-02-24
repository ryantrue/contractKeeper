package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ryantrue/contractKeeper.git/internal/config"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken создает JWT токен для аутентифицированных пользователей
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Токен истекает через 24 часа
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthenticateToken проверяет валидность JWT токена
func AuthenticateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
