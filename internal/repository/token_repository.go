package repository

import (
	"github.com/tamim447/internal/domain"
)

type TokenRepository interface {
	Save(token *domain.MagicLinkToken) error
	Find(token string) (*domain.MagicLinkToken, error)
	MarkUsed(token string) error
	InvalidateUserTokens(userID string) error
}
