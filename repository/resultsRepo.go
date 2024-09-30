package repository

import (
	"context"
	"fmt"
	"test/domain"

	"github.com/jmoiron/sqlx"
)

type ResultsDBRepo struct {
	db *sqlx.DB
}

func NewReultsDBRepo(client *sqlx.DB) *ResultsDBRepo {
	return &ResultsDBRepo{client}
}

func (r *ResultsDBRepo) Save(ctx context.Context, result domain.Result) error {
	query := `INSERT INTO results (
		id, test_user_id, percent
	)
	VALUES ($1, $2, $3)`
	queryArgs := []interface{}{
		result.ID,
		result.TestUserID,
		result.Percent,
	}

	if _, err := r.db.ExecContext(ctx, query, queryArgs...); err != nil {
		return fmt.Errorf("could not save user: %w", err)
	}
	return nil
}
