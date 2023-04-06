package repository

import (
	"context"

	"github.com/kuruyasin8/ginger/errors"
	"github.com/kuruyasin8/ginger/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	*Repository
	collection *mongo.Collection
}

var users *Users

func NewUsersRepository(ctx context.Context, repo *Repository) *Users {
	collection := repo.database.Collection("users")

	if users == nil {
		users = &Users{
			Repository: repo,
			collection: collection,
		}
	}

	return users
}

func (r *Users) GetSingleUser(ctx context.Context, filter interface{}) (*model.User, error) {
	if filter == nil {
		filter = new(bson.M)
	}

	user := new(model.User)

	if err := r.collection.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, errors.NewNotFound("user (%d) not found", filter.(bson.M)["_id"])
	}

	return user, nil
}

func (r *Users) GetMultipleUsers(ctx context.Context, filter interface{}) ([]*model.User, error) {
	if filter == nil {
		filter = new(bson.M)
	}

	users := make([]*model.User, 0)

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
