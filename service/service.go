package service

import (
	"context"

	"github.com/kuruyasin8/ginger/repository"
)

type Service struct {
	usersRepository *repository.Users
}

var service *Service

func New(ctx context.Context, repo *repository.Repository) *Service {
	if service == nil {
		service = &Service{
			usersRepository: repository.NewUsersRepository(ctx, repo),
		}
	}

	return service
}
