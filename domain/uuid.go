package domain

import "github.com/google/uuid"

type UUID string

func NewUUID() UUID {
	ID := uuid.New()
	return UUID(ID.String())
}
