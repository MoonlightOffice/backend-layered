package user

import (
	"errors"
	"giants/pkg/domain"
	"giants/pkg/repository/db"
	"giants/pkg/util"
)

type userService struct{}

func (us userService) Register(email string) (*domain.User, error) {
	uObj, ok := domain.NewUser(email)
	if !ok {
		return nil, errors.New("invalid user input")
	}

	dbrepo, err := db.NewDBRepository()
	if err != nil {
		return nil, util.ErrBuilder(err)
	}
	defer dbrepo.Close()

	err = dbrepo.AddUser(uObj)
	if err != nil {
		if errors.Is(err, db.ErrDuplicated) {
			return nil, errors.New("email already taken")
		}

		return nil, util.ErrBuilder(err)
	}

	return uObj, nil
}
