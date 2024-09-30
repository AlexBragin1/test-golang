package domain

import (
	"testing"
)

func TestShortPassword(t *testing.T) {
	err := PasswordPlain("abc").Valid()
	if err == nil {
		t.Errorf("password validation must return error")
	}
}
func TestPasswordWithoutUpperLetters(t *testing.T) {
	err := PasswordPlain("should_be_err2024").Valid()
	if err == nil {
		t.Errorf("password without upper letters must raise error")
	}
}
func TestPasswordWithoutLowerLetters(t *testing.T) {
	err := PasswordPlain("SHOULD_BE_ERR2024").Valid()
	if err == nil {
		t.Errorf("password without lower letters must raise error")
	}
}
func TestPasswordWithoutLetters(t *testing.T) {
	err := PasswordPlain("0123456789").Valid()
	if err == nil {
		t.Errorf("password without letters must raise error")
	}
}

func TestPasswordWithoutDigits(t *testing.T) {
	err := PasswordPlain("no digits here").Valid()
	if err == nil {
		t.Errorf("password validation must return error")
	}
}

func TestPasswordOK(t *testing.T) {
	err := PasswordPlain("shouldbeOK42").Valid()
	if err != nil {
		t.Errorf("password should be ok")
	}
}
func TestPasswordInvalidSymbol1(t *testing.T) {
	err := PasswordPlain("should*beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol2(t *testing.T) {
	err := PasswordPlain("should[]beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}

func TestPasswordInvalidSymbol3(t *testing.T) {
	err := PasswordPlain("should~beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol4(t *testing.T) {
	err := PasswordPlain("should&beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol5(t *testing.T) {
	err := PasswordPlain("should$beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol6(t *testing.T) {
	err := PasswordPlain("should^beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol7(t *testing.T) {
	err := PasswordPlain("test{}beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}

func TestPasswordInvalidSymbol8(t *testing.T) {
	err := PasswordPlain("should+beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}

func TestPasswordInvalidSymbol9(t *testing.T) {
	err := PasswordPlain("should-beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol10(t *testing.T) {
	err := PasswordPlain("should:beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol11(t *testing.T) {
	err := PasswordPlain("should><beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
func TestPasswordInvalidSymbol12(t *testing.T) {
	err := PasswordPlain("should.beOK42").Valid()
	if err == nil {
		t.Errorf("password contains invalid symbol")
	}
}
