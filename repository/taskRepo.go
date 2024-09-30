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

type TasksDBRepo struct {
	db *sqlx.DB
}

func NewTasksDBRepo(client *sqlx.DB) *TasksDBRepo {
	return &TasksDBRepo{client}
}

func (r *TasksDBRepo) GetByVariantID(ctx context.Context, variantID domain.UUID) (*domain.Task, error) {
	query := `SELECT id, variant_id, , description, correct_answer, options
	FROM task
	WHERE variant_id = $1 
	LIMIT 1`
	queryArgs := []interface{}{
		variantID,
	}
	var task domain.Task

	if err := r.db.GetContext(ctx, &task, query, queryArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors2.NewNotFoundError("user not found")
		} else {
			return nil, fmt.Errorf("unexpected error: %w", err)
		}
	}

	return &task, nil
}

func (r *TasksDBRepo) CountByVariantID(ctx context.Context, variantID domain.UUID) (int, error) {
	query := `SELECT count(*)
	FROM tasks
	WHERE  variant_id = $1 
	 `

	var count int
	queryArgs := []interface{}{
		variantID,
	}

	if err := r.db.GetContext(ctx, &count, query, queryArgs...); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TasksDBRepo) FindByVariantID(ctx context.Context, variantID domain.UUID) ([]domain.Task, error) {
	query := `SELECT (*)
	FROM variants
	ORDER BY ASC`

	var tasks []domain.Task
	queryArgs := []interface{}{
		variantID,
	}

	if err := r.db.SelectContext(ctx, &tasks, query, queryArgs...); err != nil {
		return nil, err
	}

	return tasks, nil
}
