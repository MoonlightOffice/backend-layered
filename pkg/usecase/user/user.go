package user

import "giants/pkg/domain"

type IUserService interface {
	Register(email string) (*domain.User, error)
}

var UserService IUserService = userService{}
