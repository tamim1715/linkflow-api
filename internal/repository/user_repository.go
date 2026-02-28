package repository

import (
	"errors"

	"github.com/tamim447/internal/domain"
)

var ErrNotFound = errors.New("user not found")

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
}
