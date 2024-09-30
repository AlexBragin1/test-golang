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

func (r *AnswersDBRepo) CountByAnswerID(ctx context.Context, answerID domain.UUID) (int, error) {
	query := `SELECT count(*)
	FROM answers
	WHERE answer_id = $1 
	`

	queryArgs := []interface{}{
		answerID,
	}
	var count int

	if err := r.db.GetContext(ctx, &count, query, queryArgs...); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *AnswersDBRepo) GetByAnswerID(ctx context.Context, AnswerID domain.UUID) (*domain.Answer, error) {
	query := `SELECT id, test_user_id, answer
	FROM answer
	WHERE test_user_id = $1 
	LIMIT 1`
	queryArgs := []interface{}{
		AnswerID,
	}
	var answer domain.Answer

	if err := r.db.GetContext(ctx, &answer, query, queryArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NewNotFoundError("user not found")
		} else {
			return nil, fmt.Errorf("unexpected error: %w", err)
		}
	}

	return &answer, nil
}

func (r *AnswersDBRepo) Update(ctx context.Context, answer domain.Answer) error {
	query := `UPDATE answers
	SET answer = $1, 
	WHERE answers_id = $2 `

	queryArgs := []interface{}{
		answer.Answer,
		answer.ID,
	}

	if _, err := r.db.ExecContext(ctx, query, queryArgs...); err != nil {
		return fmt.Errorf("could not update answer: %w", err)
	}

	return nil
}
