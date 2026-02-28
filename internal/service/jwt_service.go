package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tamim447/internal/config"
)

type JWTService struct {
	Secret string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		Secret: secret,
	}
}

func (j *JWTService) Generate(userID string) (string, error) {

	expireHour, err := strconv.Atoi(config.JWTExpireHours)
	if err != nil {
		expireHour = 7 * 24
	}

	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Duration(expireHour) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.Secret))
}

func (j *JWTService) Validate(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	userID := claims["userId"].(string)

	return userID, nil
}
