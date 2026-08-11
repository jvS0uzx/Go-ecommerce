package service

import (
	"ecommerce-api/internal/model"
	"ecommerce-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService {
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)


	if err != nil{
		return err
	}

	user.PasswordHash = string(hashedBytes)

	return s.repo.CreateUser(user)

}