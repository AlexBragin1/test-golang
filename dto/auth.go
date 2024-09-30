package dto

import (
	"test/domain"
	"test/errors"
)

type RegisterReq struct {
	Login           domain.Login         `json:"login"`
	Password        domain.PasswordPlain `json:"password"`
	PasswordConfirm domain.PasswordPlain `json:"password_confirm"`
}

func (r *RegisterReq) Validate() error {
	if len(r.Login) < 5 {
		return errors.NewValidationError("password too short")
	}

	if len(r.Password) < 3 {
		return errors.NewValidationError("password too short")
	}

	if r.Password != r.PasswordConfirm {
		return errors.NewValidationError("passwords do not match")
	}

	return nil
}

type RegisterRes struct {
	Token string `json:"token"`
}

type LoginReq struct {
	Login    domain.Login         `json:"login"`
	Password domain.PasswordPlain `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

func (r *LoginReq) Validate() error {
	if len(r.Login) < 5 {
		return errors.NewValidationError("password too short")
	}

	if len(r.Password) < 3 {
		return errors.NewValidationError("password too short")
	}

	return nil
}

type LoginOutReq struct {
	UserID domain.UUID `json:"UserID"`
}

type LoginOutRes struct{}

func (r *LoginOutReq) Validate() error {
	if len(r.UserID) > 0 {
		return errors.NewValidationError("invalid user id ")
	}

	return nil
}
