package user

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Email    string
	Password string
}
type User struct {
	ID    string
	Email string
}

type ErrorsInscription struct {
	Email    string
	Password string
}

func (u *UserCreate) Validate(password string) *ErrorsInscription {
	m := ErrorsInscription{}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		m.Email = "Address email invalid"
	}
	if u.Password != password {
		m.Password = "password do not match"
	}
	if m != (ErrorsInscription{}) {
		return &m
	}
	return nil
}
func (u *UserCreate) ValidatePassword(password string) error {
	var specialRunes = "!@#$%^&*"
	m := map[string]bool{
		"special":   false,
		"minuscule": false,
		"majuscule": false,
		"numbers":   false,
	}
	if len(password) < 8 {
		return errors.New("minimum 8 characters requis")
	}
	for _, r := range password {
		switch {
		case unicode.IsPunct(r) || unicode.IsSymbol(r) || strings.ContainsRune(specialRunes, r):
			m["special"] = true
		case unicode.IsLower(r):
			m["minuscule"] = true
		case unicode.IsUpper(r):
			m["majuscule"] = true
		case unicode.IsDigit(r):
			m["numbers"] = true
		}
	}
	for i := 0; i < len(m); i++ {
		switch {
		case !m["special"]:
			return fmt.Errorf("minimum un special character : (%s)", specialRunes)
		case !m["minuscule"]:
			return errors.New("minimum une lettre minuscule")
		case !m["majuscule"]:
			return errors.New("minimum une lettre majuscule")
		case !m["numbers"]:
			return errors.New("minimum un chiffre")
		}

	}
	return nil
}
func (u *UserCreate) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
