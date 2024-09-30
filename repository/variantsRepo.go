package repository

import (
	"context"
	"test/domain"

	"github.com/jmoiron/sqlx"
)

type VariantsDBRepo struct {
	db *sqlx.DB
}

func NewVariantsDBRepo(client *sqlx.DB) *VariantsDBRepo {
	return &VariantsDBRepo{client}
}
func (r *VariantsDBRepo) Save(ctx context.Context) {

	return
}

func (r *VariantsDBRepo) FindAll(ctx context.Context) ([]domain.Variant, error) {
	query := `SELECT (*)
	FROM variants
	ORDER BY `

	var variants []domain.Variant

	if err := r.db.SelectContext(ctx, &variants, query); err != nil {
		return nil, err
	}

	return variants, nil
}
