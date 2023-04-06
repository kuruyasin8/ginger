package service

import (
	"context"

	"github.com/kuruyasin8/ginger/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserQuery struct {
	ID      uint   `param:"uid"`
	Page    uint   `query:"page"`
	PerPage uint   `query:"per_page"`
	Filter  string `query:"filter"`
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
