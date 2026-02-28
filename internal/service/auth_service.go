package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/tamim447/internal/config"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/email"
	"github.com/tamim447/internal/repository"
)

type AuthService struct {
	Users  repository.UserRepository
	Tokens repository.TokenRepository
	Email  email.Sender

	TokenGen *TokenService
	JWT      *JWTService
}

func NewAuthService(
	users repository.UserRepository,
	tokens repository.TokenRepository,
	tokenGen *TokenService,
	jwt *JWTService,
	email email.Sender,
) *AuthService {

	return &AuthService{
		Users:    users,
		Tokens:   tokens,
		TokenGen: tokenGen,
		JWT:      jwt,
		Email:    email,
	}
}

func (s *AuthService) RequestMagicLink(emailAddr string) error {

	if emailAddr == "" {
		return errors.New(constants.InvalidEmailAddress)
	}

	// 1. Find user
	user, err := s.Users.FindByEmail(emailAddr)

	// 2. Auto sign-up if not found
	if errors.Is(err, repository.ErrNotFound) {
		user = &domain.User{Email: emailAddr}
		if err := s.Users.Create(user); err != nil {
			return err
		}
	}

	// 3. Generate token
	token := s.TokenGen.Generate(user.ID)

	// Invalidate previous unused tokens
	_ = s.Tokens.InvalidateUserTokens(user.ID)

	// 4. Save token
	if err := s.Tokens.Save(token); err != nil {
		return err
	}

	// 5. Send magic link (mocked)
	link := fmt.Sprintf(config.VerifyTokenURI+"?token=%s", token.Token)

	return s.Email.SendMagicLink(emailAddr, link)
}

func (s *AuthService) Verify(token string) (string, error) {

	if token == "" {
		return "", errors.New(constants.InvalidRequest)
	}

	magic, err := s.Tokens.Find(token)
	if err != nil {
		return "", err
	}

	if magic.Used {
		return "", errors.New(constants.TokenAlreadyUsed)
	}

	if time.Now().After(magic.ExpiresAt) {
		return "", errors.New(constants.TokenExpired)
	}

	if err := s.Tokens.MarkUsed(token); err != nil {
		return "", err
	}

	// Generate JWT userID
	return s.JWT.Generate(magic.UserID)
}
