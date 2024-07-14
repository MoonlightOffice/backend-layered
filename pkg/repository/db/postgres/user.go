package postgres

import (
	"context"
	"errors"
	"time"

	"giants/pkg/domain"
	"giants/pkg/util"

	"github.com/jackc/pgx/v5"
)

func (p Postgres) AddUser(uObj *domain.User) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (user_id, email, created_at) VALUES ($1, $2, $3)`
	_, err := p.crud.Exec(
		ctx,
		stmt,
		uObj.UserId,
		uObj.Email,
		uObj.CreatedAt,
	)
	if err != nil {
		if isErrDuplicate(err) {
			return util.ErrBuilder(ErrDuplicated)
		}

		return util.ErrBuilder(err)
	}

	return nil
}

func (p Postgres) FindUserById(userId string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var (
		email     string
		createdAt time.Time
	)

	stmt := `SELECT email, created_at FROM users WHERE user_id = $1 FOR UPDATE`
	err := p.crud.QueryRow(ctx, stmt, userId).Scan(&email, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, util.ErrBuilder(ErrNotExist)
		}

		return nil, util.ErrBuilder(err)
	}

	return &domain.User{
		UserId:    userId,
		Email:     email,
		CreatedAt: createdAt,
	}, nil
}

func (p Postgres) FindUserByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	var (
		userId    string
		createdAt time.Time
	)

	stmt := `SELECT user_id, created_at FROM users WHERE email = $1 FOR UPDATE`
	err := p.crud.QueryRow(ctx, stmt, email).Scan(&userId, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, util.ErrBuilder(ErrNotExist)
		}

		return nil, util.ErrBuilder(err)
	}

	return &domain.User{
		UserId:    userId,
		Email:     email,
		CreatedAt: createdAt,
	}, nil
}
