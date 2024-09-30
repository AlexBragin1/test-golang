package domain

type Task struct {
	ID            UUID   `db:"id" json:"id"`
	VariantID     int    `db:"variant_id" json:"variant_id"`
	Description   string `db:"description" json:"description"`
	CorrectAnswer string `db:"correct_answer" json:"-"`
	Options       string `db:"options" db: json:"options"`
}
