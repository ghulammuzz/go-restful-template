package service

import (
	"github.com/ghulammuzz/go-restful-template/internal/errors"
	"github.com/ghulammuzz/go-restful-template/internal/model"
	"github.com/ghulammuzz/go-restful-template/internal/repository"
	"github.com/ghulammuzz/go-restful-template/pkg/utils"
)

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) RegisterUser(username, password string) error {
	if username == "" {
		return errors.NewAppError(errors.ErrUsernameRequired, "Username is required")
	}
	if password == "" {
		return errors.NewAppError(errors.ErrPasswordRequired, "Password is required")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.NewAppError(errors.ErrHashingFailed, "Error hashing password")
	}

	user := &model.User{
		Username: username,
		Password: hashedPassword,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "Unable to register user")
	}

	return nil
}
