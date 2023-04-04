package main

import (
	"context"

	"github.com/kuruyasin8/ginger/repository"
)

type User struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

func main() {
	ctx := context.Background()

	repo := repository.New(ctx)
	if err := repo.Connect(ctx); err != nil {
		panic(err)
	}
	defer repo.Close(ctx)

	usersRepository := repository.NewUsersRepository(ctx, repo)

	user, err := usersRepository.GetSingleUser(ctx, nil)
	if err != nil {
		panic(err)
	}

	println(user.ID)
	println(user.Name)
}
