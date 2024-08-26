package repository

import "github.com/ghulammuzz/go-restful-template/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
}
