package mongodb

import (
	"context"
	"errors"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/domain"
	"github.com/tamim447/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrNotFound = errors.New("not found")

//type TokenRepo struct {
//	Collection *mongodb.Collection
//}
//
//func (r *TokenRepo) Save(token *domain.MagicLinkToken) error {
//	_, err := r.Collection.InsertOne(context.TODO(), token)
//	return err
//}
//
//func (r *TokenRepo) Find(token string) (*domain.MagicLinkToken, error) {
//
//	var doc domain.MagicLinkToken
//
//	err := r.Collection.FindOne(context.TODO(), map[string]any{
//		"token": token,
//	}).Decode(&doc)
//
//	if err == mongodb.ErrNoDocuments {
//		return nil, repository.ErrNotFound
//	}
//
//	return &doc, err
//}
//
//func (r *TokenRepo) MarkUsed(token string) error {
//	_, err := r.Collection.UpdateOne(
//		context.TODO(),
//		map[string]any{"token": token},
//		map[string]any{
//			"$set": map[string]any{"used": true},
//		},
//	)
//
//	return err
//}

type TokenRepo struct {
	Collection *mongo.Collection
}

func NewMongoTokenRepository(db *mongo.Database) repository.TokenRepository {
	return &TokenRepo{
		Collection: db.Collection("magic_link_tokens"),
	}
}

func (r *TokenRepo) Save(token *domain.MagicLinkToken) error {
	_, err := r.Collection.InsertOne(context.TODO(), token)
	return err
}

func (r *TokenRepo) Find(token string) (*domain.MagicLinkToken, error) {
	var doc domain.MagicLinkToken

	err := r.Collection.FindOne(context.TODO(), map[string]any{
		constants.Token: token,
	}).Decode(&doc)

	if err == mongo.ErrNoDocuments {
		return nil, ErrNotFound
	}

	return &doc, err
}

func (r *TokenRepo) MarkUsed(token string) error {
	_, err := r.Collection.UpdateOne(
		context.TODO(),
		map[string]any{constants.Token: token},
		map[string]any{
			"$set": map[string]any{"used": true},
		},
	)

	return err
}
