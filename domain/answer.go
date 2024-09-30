package domain

type Answer struct {
	ID         UUID   `db:"id" json:"id"`
	TestUserID UUID   `db:"test_user_id" json:"test_user_id"`
	Answer     string `db:"answer" json:"answer"`
}

func NewAnswer(testUserId UUID, answer string) *Answer {

	return &Answer{
		ID:         NewUUID(),
		TestUserID: testUserId,
		Answer:     answer,
	}
}
