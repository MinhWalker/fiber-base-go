package services

import (
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/repository"
	"github.com/pkg/errors"
)

type UserService interface {
	UpsertUser(user *model.User) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) UpsertUser(user *model.User) (*model.User, error) {
	var userDB *model.User
	var err error
	if userDB, err = s.userRepo.UpsertUser(user); err != nil {
		return nil, errors.Wrap(err, "studentService.CreateStudent")
	}

	return userDB, nil
}
