package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/service"
	"testing"
	"time"
)

type FakeTokenRepo struct {
	Token *domain.MagicLinkToken
}

func (f *FakeTokenRepo) Save(token *domain.MagicLinkToken) error {
	// Not needed in this test
	return nil
}

func (f *FakeTokenRepo) Find(token string) (*domain.MagicLinkToken, error) {
	return f.Token, nil
}

func (f *FakeTokenRepo) MarkUsed(token string) error {
	f.Token.Used = true
	return nil
}

func TestVerifyMarksTokenUsed(t *testing.T) {

	repo := &FakeTokenRepo{
		Token: &domain.MagicLinkToken{
			Token:     "abc",
			UserID:    "user123",
			ExpiresAt: time.Now().Add(10 * time.Minute),
			Used:      false,
		},
	}

	auth := &service.AuthService{
		Tokens: repo,
		JWT:    &service.JWTService{Secret: "test"},
	}

	_, err := auth.Verify("abc")

	assert.NoError(t, err)
	assert.True(t, repo.Token.Used)
}
