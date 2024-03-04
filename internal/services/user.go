package services

import "github.com/Rizabekus/registration-api/internal/models"

type UserService struct {
	storage models.UserStorage
}

func CreateUserService(storage models.UserStorage) *UserService {
	return &UserService{storage: storage}
}
