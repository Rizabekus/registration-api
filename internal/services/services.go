package services

import (
	"github.com/Rizabekus/registration-api/internal/models"
	"github.com/Rizabekus/registration-api/internal/storage"
)

type Services struct {
	UserService models.UserService
}

func ServiceInstance(storage *storage.Storage) *Services {
	return &Services{
		UserService: CreateUserService(storage.UserStorage),
	}
}
