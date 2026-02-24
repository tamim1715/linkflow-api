package service

import (
	"github.com/google/uuid"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/repository"
	"github.com/tamim447/internal/slack"
	"time"
)

type FeedbackService struct {
	Repo  repository.FeedbackRepository
	Slack slack.Client
}

func (s *FeedbackService) Submit(userID string, message string) error {

	feedback := &domain.Feedback{
		ID:        uuid.NewString(),
		UserID:    userID,
		Message:   message,
		CreatedAt: time.Now(),
	}

	if err := s.Repo.Save(feedback); err != nil {
		return err
	}

	return s.Slack.Publish("New feedback from " + userID + ": " + message)
}
