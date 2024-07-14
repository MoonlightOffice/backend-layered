package domain

import (
	"time"

	"giants/pkg/util"
)

type User struct {
	UserId    string
	Email     string
	CreatedAt time.Time
}

func NewUser(email string) (*User, bool) {
	if len(email) == 0 {
		return nil, false
	}

	return &User{
		UserId:    util.NewID(),
		Email:     email,
		CreatedAt: time.Now(),
	}, true
}
