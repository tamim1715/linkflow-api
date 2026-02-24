package mongodb

import (
	"context"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeedbackRepo struct {
	Collection *mongo.Collection
}

func NewMongoFeedbackRepository(db *mongo.Database) repository.FeedbackRepository {
	return &FeedbackRepo{
		Collection: db.Collection("feedbacks"),
	}
}

func (r *FeedbackRepo) Save(feedback *domain.Feedback) error {
	_, err := r.Collection.InsertOne(context.TODO(), feedback)
	return err
}
