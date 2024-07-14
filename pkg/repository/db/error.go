package db

import "giants/pkg/repository/db/postgres"

var (
	ErrNotExist   = postgres.ErrNotExist
	ErrDuplicated = postgres.ErrDuplicated
)
