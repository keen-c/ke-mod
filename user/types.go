package user

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode"
)

type User struct {
	id       string
	email    string
	password string
}

func (u *User) Validate(password string) map[string]string {
	m := make(map[string]string)
	if _, err := mail.ParseAddress(u.email); err != nil {
		m["email"] = "Address email invalid"
	}
	if u.password != password {
		m["password"] = "password do not match"
	}
	if len(m) > 0 {
		return m
	}
	return nil
}
func (u *User) ValidatePassword(password string) error {
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
	for _, r := range password{
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
	for i := 0 ; i < len(m); i++ {
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
