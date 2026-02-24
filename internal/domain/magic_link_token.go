package domain

import "time"

type MagicLinkToken struct {
	Token     string
	UserID    string
	ExpiresAt time.Time
	Used      bool
}
