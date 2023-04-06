package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/kuruyasin8/ginger/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserQuery struct {
	ID      uint   `param:"uid"`
	Page    uint   `query:"page"`
	PerPage uint   `query:"per_page"`
	Filter  string `query:"filter"`
}

func (s *Service) Register(ctx context.Context, user *model.User) error {
	salt := make([]byte, 32)
	rand.Read(salt)

	raw := append([]byte(user.Password), salt...)
	hash := sha256.Sum256(raw)

	encodedHash := base64.StdEncoding.EncodeToString(hash[:])
	encodedSalt := base64.StdEncoding.EncodeToString(salt[:])

	user.Credentials = new(model.Credentials)
	user.Credentials.Hash = encodedHash
	user.Credentials.Salt = encodedSalt
	user.Credentials.Verified = true
	now := time.Now().UnixMilli()
	user.Credentials.CreatedAt = now
	user.Credentials.ModifiedAt = now

	if err := s.usersRepository.InsertSingleUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetSingleUser(ctx context.Context, query *UserQuery) (*model.User, error) {
	filter := bson.M{"_id": query.ID}

	user, err := s.usersRepository.GetSingleUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetMultipleUsers(ctx context.Context, query *UserQuery) ([]*model.User, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query.Filter, "$options": "i"}},
		},
	}

	users, err := s.usersRepository.GetMultipleUsers(ctx, filter)
	if err != nil {
		return nil, err
	}

	if len(users) < int(query.Page*query.PerPage) {
		return users, nil
	}

	users = users[(query.Page-1)*query.PerPage : query.PerPage*query.Page]

	return users, nil
}
