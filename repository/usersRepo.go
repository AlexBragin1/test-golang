package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test/domain"
	errors2 "test/errors"

	"github.com/jmoiron/sqlx"
)

type UsersDBRepo struct {
	db *sqlx.DB
}

func NewUsersDBRepo(client *sqlx.DB) *UsersDBRepo {
	return &UsersDBRepo{client}
}

func (r *UsersDBRepo) Save(ctx context.Context, user domain.User) error {
	query := `INSERT
			INTO users(id, login, password, groups, auth,start_session_at,end_session_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	queryArgs := []interface{}{
		user.ID,
		user.Login,
		user.Password,
		user.Groups,
		user.Auth,
		user.StartSessionAt,
		user.EndSessionAt,
	}
	if _, err := r.db.ExecContext(ctx, query, queryArgs...); err != nil {

		return fmt.Errorf("could not save user: %w", err)
	}

	return nil
}

func (r *UsersDBRepo) FindByLogin(ctx context.Context, login domain.Login) (*domain.User, error) {
	query := `SELECT id, login, password, groups, auth, group,start_session_at,end_session_at
	FROM users
	WHERE login = $1 
	LIMIT 1`

	var user domain.User

	if err := r.db.Get(&user, query, login); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NewNotFoundError("user not found")
		} else {
			return nil, fmt.Errorf("unexpected error: %w", err)
		}
	}

	return &user, nil
}

func (r *UsersDBRepo) FindByID(ctx context.Context, userID domain.UUID) (*domain.User, error) {
	query := `SELECT id, login, password, groups, auth, group,start_session_at,end_session_at
	FROM users
	WHERE id = $1 
	LIMIT 1`

	var user domain.User

	if err := r.db.Get(&user, query, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NewNotFoundError("user not found")
		} else {
			return nil, fmt.Errorf("unexpected error: %w", err)
		}
	}

	return &user, nil
}

func (r *UsersDBRepo) Update(ctx context.Context, user domain.User) error {
	query := `UPDATE users
	SET  auth = $1, start_session_at=$2, end_session_at =$3
	WHERE login = $4 `

	queryArgs := []interface{}{
		user.Auth,
		user.StartSessionAt,
		user.EndSessionAt,
		user.Login,
	}

	if _, err := r.db.ExecContext(ctx, query, queryArgs...); err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	return nil
}

func (r *UsersDBRepo) CountByLogin(ctx context.Context, login domain.Login) (int, error) {
	query := `SELECT count(*)
	FROM users
	WHERE login = $1 `

	queryArgs := []interface{}{
		login,
	}
	var count int

	if err := r.db.GetContext(ctx, &count, query, queryArgs...); err != nil {
		return 0, err
	}

	return count, nil
}
