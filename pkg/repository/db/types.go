package db

import "giants/pkg/domain"

type IDB interface {
	// Close (or release) db connection
	Close()

	/* User */

	// err: ErrDuplicated
	AddUser(uObj *domain.User) error

	// err: ErrNotExist
	FindUserById(userId string) (*domain.User, error)

	// err: ErrNotExist
	FindUserByEmail(email string) (*domain.User, error)

	/* Article */

	/* Payment */

	/* Post */
}
