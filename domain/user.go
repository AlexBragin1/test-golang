package domain

import (
	"test/domain/groups"
	"time"
)

type User struct {
	ID             UUID          `db:"id" json:"id"`
	Login          Login         `db:"login" json:"login"`
	Password       PasswordHash  `db:"password" json:"-"`
	Groups         groups.Groups `db:"groups" json:"groups"`
	Auth           bool          `db:"auth" json:auth`
	StartSessionAt time.Time     `db:"start_session_at" json:"strat_sesion_at"`
	EndSessionAt   time.Time     `db:"end_session_at" json:"end_sesion_at"`
}

func NewUser(login Login, password PasswordPlain, group groups.Groups) (*User, error) {
	now := time.Now()

	hash, err := password.Hash()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:             NewUUID(),
		Login:          login,
		Password:       hash,
		Groups:         groups.User,
		Auth:           false,
		StartSessionAt: now,
		EndSessionAt:   now,
	}, nil
}

func (u User) Update(startSessionAt time.Time, endSessionAt time.Time) *User {
	return &User{
		ID:             u.ID,
		Login:          u.Login,
		Password:       u.Password,
		Groups:         u.Groups,
		Auth:           false,
		StartSessionAt: startSessionAt,
		EndSessionAt:   endSessionAt,
	}
}
