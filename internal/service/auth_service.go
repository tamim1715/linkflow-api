package service

import (
	"errors"
	"fmt"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/email"
	"github.com/tamim447/internal/repository"
	"time"
)

type AuthService struct {
	Users  repository.UserRepository
	Tokens repository.TokenRepository
	Email  email.Sender

	TokenGen *TokenService
	JWT      *JWTService
}

func (s *AuthService) RequestMagicLink(emailAddr string) error {

	// 1. Find user
	user, err := s.Users.FindByEmail(emailAddr)

	// 2. Auto sign-up if not found
	if err == repository.ErrNotFound {
		user = &domain.User{Email: emailAddr}
		if err := s.Users.Create(user); err != nil {
			return err
		}
	}

	// 3. Generate token
	token := s.TokenGen.Generate(user.ID)

	// Invalidate previous unused tokens
	s.Tokens.InvalidateUserTokens(user.ID)

	// 4. Save token
	if err := s.Tokens.Save(token); err != nil {
		return err
	}

	// 5. Send magic link (mocked)
	link := fmt.Sprintf("myapp://api/auth/verify?token=%s", token.Token)

	return s.Email.SendMagicLink(emailAddr, link)
}

func (s *AuthService) Verify(token string) (string, error) {

	magic, err := s.Tokens.Find(token)
	if err != nil {
		return "", err
	}

	if magic.Used {
		return "", errors.New("token already used")
	}

	if time.Now().After(magic.ExpiresAt) {
		return "", errors.New("token expired")
	}

	if err := s.Tokens.MarkUsed(token); err != nil {
		return "", err
	}

	// Generate JWT userID
	return s.JWT.Generate(magic.UserID)
}
