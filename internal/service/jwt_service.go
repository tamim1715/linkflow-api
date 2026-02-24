package service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//type JWTService struct {
//	Secret string
//}
//
//func (j *JWTService) Generate(userID string) (string, error) {
//
//	claims := jwt.MapClaims{
//		"userId": userID,
//		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	return token.SignedString([]byte(j.Secret))
//}

type JWTService struct {
	Secret string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		Secret: secret,
	}
}

func (j *JWTService) Generate(userID string) (string, error) {

	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
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
