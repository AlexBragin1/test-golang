package repository

import (
	"context"
	"fmt"
	"test/domain"

	"github.com/jmoiron/sqlx"
)

type AnswersDBRepo struct {
	db *sqlx.DB
}

func NewAnswersDBRepo(client *sqlx.DB) *AnswersDBRepo {

	return &AnswersDBRepo{client}
}
func (r *AnswersDBRepo) Save(ctx context.Context, answers domain.Answer) error {
	query := `INSERT INTO results (
		id, test_user_id, answer
	)
		VALUES ($1, $2, $3)`

	queryArgs := []interface{}{}

	if _, err := r.db.ExecContext(ctx, query, queryArgs...); err != nil {
		return fmt.Errorf("could not save answers: %w", err)
	}

	return nil
}
