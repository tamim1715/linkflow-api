package mongodb

import (
	"context"
	"github.com/google/uuid"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) repository.UserRepository {
	return &UserRepo{
		Collection: db.Collection(constants.UserCollection),
	}
}

func (r *UserRepo) FindByEmail(email string) (*domain.User, error) {
	var doc struct {
		ID    string `bson:"_id"`
		Email string `bson:"email"`
	}

	err := r.Collection.FindOne(context.TODO(), map[string]string{
		"email": email,
	}).Decode(&doc)

	if err == mongo.ErrNoDocuments {
		return nil, repository.ErrNotFound
	}

	return &domain.User{
		ID:    doc.ID,
		Email: doc.Email,
	}, nil
}

func (r *UserRepo) Create(user *domain.User) error {
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()

	_, err := r.Collection.InsertOne(context.TODO(), map[string]any{
		"_id":       user.ID,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})

	return err
}
