package domain

import (
	"regexp"

	"test/errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordPlain string
type PasswordHash string

var ContainsUpperLetter = regexp.MustCompile(`[A-Z]+`).MatchString
var ContainsLowerLetter = regexp.MustCompile(`[a-z]+`).MatchString
var ContainsDigit = regexp.MustCompile(`[0-9]+`).MatchString
var ContainsInvalidSymbol = regexp.MustCompile(`[\~!?@#$%\^&*_\-+()\[\]\{}><\/\\|"'.,:]+`).MatchString

func (p PasswordPlain) Valid() *errors.AppError {
	if len(p) < 8 {
		return errors.NewValidationError("password is too short")
	}

	if !ContainsLowerLetter(string(p)) {
		return errors.NewValidationError("password must contain lower letter")
	}

	if !ContainsUpperLetter(string(p)) {
		return errors.NewValidationError("password must contain upper letter")
	}

	if !ContainsDigit(string(p)) {
		return errors.NewValidationError("password must contain digit")
	}

	if ContainsInvalidSymbol(string(p)) {
		return errors.NewValidationError("password contains invalid symbol")
	}

	return nil
}

func (p PasswordPlain) Hash() (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	return PasswordHash(bytes), err
}

func (p PasswordPlain) CheckWithHash(h PasswordHash) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(h))

	return err == nil
}

func (h PasswordHash) CheckWithPlain(p PasswordPlain) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))

	return err == nil
}
