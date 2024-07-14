package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrNotExist   = errors.New("record does not exist")
	ErrDuplicated = errors.New("duplicated error")
)

func isErrDuplicate(err error) bool {
	const PgErrCodeDuplicate = "23505"

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == PgErrCodeDuplicate {
		return true
	}

	return false
}
