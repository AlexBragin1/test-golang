package domain

import (
	"time"
)

type TestUser struct {
	ID        UUID       `db:"id" json:"id"`
	UserID    UUID       `db:"user_id" json":"user_id"`
	VariantID UUID       `db:"variant_id" json:"variant_id"`
	StartAt   time.Time  `db:"start_at" json:"start_at"`
	EndAt     *time.Time `db:"end_at" json:"end_at"`
}

func NewTestUser(userID UUID, variantID UUID) *TestUser {
	now := time.Now()

	return &TestUser{
		ID:        NewUUID(),
		UserID:    userID,
		VariantID: variantID,
		StartAt:   now,
		EndAt:     &time.Time{},
	}
}
