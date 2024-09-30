package domain

import "test/errors"

type Flow string

const FLOW_REGISTRATION Flow = "REG"
const FLOW_AUTHORIZATION Flow = "AUTH"

type Login string

func (l Login) Valid() *errors.AppError {
	if len(l) < 8 {
		return errors.NewValidationError("login is too short")
	}

	if !ContainsLowerLetter(string(l)) {
		return errors.NewValidationError("login must contain lower letter")
	}

	if ContainsInvalidSymbol(string(l)) {
		return errors.NewValidationError("password contains invalid symbol")
	}

	return nil
}
