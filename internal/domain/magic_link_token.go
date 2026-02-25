package domain

import "time"

type MagicLinkToken struct {
	Token     string    `bson:"token"`
	UserID    string    `bson:"userId"`
	ExpiresAt time.Time `bson:"expiresAt"`
	Used      bool      `bson:"used"`
}
