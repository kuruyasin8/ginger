package repository

import (
	"context"

	"github.com/kuruyasin8/ginger/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client   *mongo.Client
	database *mongo.Database
}

var repository *Repository

func New(ctx context.Context) *Repository {
	if repository == nil {
		repository = &Repository{}
	}

	return repository
}

func (r *Repository) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUri))
	if err != nil {
		return err
	}

	database := client.Database("ginger")

	r.client = client
	r.database = database
	return nil
}

func (r *Repository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

func (r *Repository) Ping(ctx context.Context) error {
	return r.client.Ping(ctx, nil)
}

func (r *Repository) Database(name string) *Repository {
	r.database = r.client.Database(name)
	return r
}
