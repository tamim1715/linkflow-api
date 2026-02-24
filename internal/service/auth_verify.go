package service

import (
	"errors"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenUsed    = errors.New("token already used")
)

type VerifyResult struct {
	JWT string
}

//func (s *AuthService) VerifyMagicLink(token string) (*VerifyResult, error) {
//
//	// 1. Find token
//	stored, err := s.Tokens.Find(token)
//	if err != nil {
//		return nil, err
//	}
//
//	// 2. Check expiry
//	if time.Now().After(stored.ExpiresAt) {
//		return nil, ErrTokenExpired
//	}
//
//	// 3. Check single-use
//	if stored.Used {
//		return nil, ErrTokenUsed
//	}
//
//	// 4. Mark token as used
//	if err := s.Tokens.MarkUsed(token); err != nil {
//		return nil, err
//	}
//
//	// 5. Generate JWT session
//	jwtToken, err := s.JWT.Generate(stored.UserID)
//	if err != nil {
//		return nil, err
//	}
//
//	return &VerifyResult{JWT: jwtToken}, nil
//}
