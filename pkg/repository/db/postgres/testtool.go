package postgres

import (
	"context"
	"fmt"
	"time"
)

func DeleteAll() error {
	// Disable this function in production
	//if ... {
	//	return nil
	//}

	// Prepare DB connection
	dbrepo, err := NewPostgres()
	if err != nil {
		return fmt.Errorf("crdb.TiDBRepository.DeleteAll(): %w", err)
	}
	dbrepo.Close() // Do not defer

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// Define tables to delete data from
	tables := []string{
		"users",
	}

	for _, table := range tables {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE true`, table)
		_, err := pool.Exec(ctx, stmt)
		if err != nil {
			return err
		}
	}

	return nil
}
