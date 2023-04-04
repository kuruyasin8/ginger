package service

import (
	"context"

	"github.com/kuruyasin8/ginger/repository"
)

type Service struct {
	usersRepository *repository.Users
}

func New(ctx context.Context, repo *repository.Repository) *Service {
	return &Service{
		usersRepository: repository.NewUsersRepository(ctx, repo),
	}
}
