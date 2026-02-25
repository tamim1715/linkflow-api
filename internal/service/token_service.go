package service

import (
	"github.com/google/uuid"
	"github.com/tamim447/internal/domain"
	"time"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (t *TokenService) Generate(userID string) *domain.MagicLinkToken {
	return &domain.MagicLinkToken{
		Token:     uuid.NewString(),
		UserID:    userID,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		Used:      false,
	}
}
