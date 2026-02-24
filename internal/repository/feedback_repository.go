package repository

import "github.com/tamim447/internal/domain"

type FeedbackRepository interface {
	Save(feedback *domain.Feedback) error
}
