package domain

import "time"

type Feedback struct {
	ID        string
	UserID    string
	Message   string
	CreatedAt time.Time
}
