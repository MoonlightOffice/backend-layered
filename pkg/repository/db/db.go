package db

import "giants/pkg/repository/db/postgres"

// Do not forget to defer .Close()
func NewDBRepository() (IDB, error) {
	return postgres.NewPostgres()
}
