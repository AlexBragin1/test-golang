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

type TestsUserDBRepo struct {
	db *sqlx.DB
}

func NewTestsUserDBRepo(client *sqlx.DB) *TestsUserDBRepo {
	return &TestsUserDBRepo{client}
}

func (r *TestsUserDBRepo) Save(ctx context.Context, test domain.TestUser) error {
	query := `INSERT
			INTO tests_users(id, user_id, variant_id, start_at, end_at)
			VALUES ($1, $2, $3, $4, $5)`
	queryArgs := []interface{}{
		test.ID,
		test.UserID,
		test.VariantID,
		test.StartAt,
		test.EndAt,
	}
	_, err := r.db.ExecContext(ctx, query, queryArgs...)

	if err != nil {
		return fmt.Errorf("could not save user: %w", err)
	}

	return nil
}

func (r *TestsUserDBRepo) Update(ctx context.Context, test domain.TestUser) error {
	query := `UPDATE tests_users
	SET  start_at=$1, end_at=$2 
	WHERE user_id =$3 
	AND variant_id = $4 `
	queryArgs := []interface{}{
		test.StartAt,
		test.EndAt,
		test.UserID,
		test.VariantID,
	}

	if _, err := r.db.ExecContext(ctx, query, queryArgs...); err != nil {
		return fmt.Errorf("could not update test_user: %w", err)
	}
	return nil
}

func (r *TestsUserDBRepo) CountByVariantID(ctx context.Context, userID domain.UUID, variantID domain.UUID) (int, error) {
	query := `SELECT count(*)
	FROM tests_users
	WHERE user_id 
	AND variant_id = $1 
	 `

	var count int
	queryArgs := []interface{}{
		userID,
		variantID,
	}

	if err := r.db.GetContext(ctx, &count, query, queryArgs...); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TestsUserDBRepo) GetByVariantID(ctx context.Context, variantID domain.UUID) (*domain.TestUser, error) {
	query := `SELECT id, user_id, variant_id, start_at, end_at
	FROM tests_users
	WHERE variant_id = $1 
	LIMIT 1`
	queryArgs := []interface{}{
		variantID,
	}
	var test domain.TestUser

	if err := r.db.GetContext(ctx, &test, query, queryArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NewNotFoundError("user not found")
		} else {
			return nil, fmt.Errorf("unexpected error: %w", err)
		}
	}

	return &test, nil
}
