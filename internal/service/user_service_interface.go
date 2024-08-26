package service

type UserService interface {
	RegisterUser(username, password string) error
}
