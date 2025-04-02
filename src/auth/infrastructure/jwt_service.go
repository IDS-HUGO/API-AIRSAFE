package infrastructure

import (
	"apiusersafe/src/auth/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTServiceImpl struct {
	SecretKey string
}

func NewJWTService(secretKey string) domain.JWTService {
	return &JWTServiceImpl{SecretKey: secretKey}
}

func (s *JWTServiceImpl) GenerateToken(usuario string) (string, error) {
	claims := jwt.MapClaims{
		"usuario": usuario,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.SecretKey))
}

func (s *JWTServiceImpl) ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if usuario, ok := claims["usuario"].(string); ok {
			return usuario, nil
		}
	}

	return "", fmt.Errorf("invalid token")
}
