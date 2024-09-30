package domain

type Result struct {
	ID         UUID   `db:"id" json:"id"`
	TestUserID UUID   `db:"test__user_id" json:"test_user_id"`
	Percent    string `db:"percent" json:"percent"`
}

func NewResult(testUserId UUID, percent string) *Result {

	return &Result{
		ID:         NewUUID(),
		TestUserID: testUserId,
		Percent:    percent,
	}
}
