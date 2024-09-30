package domain

type Variant struct {
	ID     UUID   `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Number int    `db:"name" json:"number"`
}
