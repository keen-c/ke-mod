package user

import (
	"testing"
)

func Test_user_Validate(t *testing.T) {
	t.Run("validate 1", func(t *testing.T) {
		user := User{id: "1", email: "text@example.com", password: "password"}
		got := user.Validate("password")
		if got != nil {
			t.Errorf("%v", got)
		}
	})
	t.Run("validate 2", func(t *testing.T) {
		user := User{id: "1", email: "textexample.com", password: "password"}
		got := user.Validate("password")
		if len(got) == 0 {
			t.Errorf("%v", got)
		}
	})
}

func Test_user_ValidatePassword(t *testing.T) {
	t.Run("validate password 1", func(t *testing.T) {
		var user User
		got := user.ValidatePassword("un11mot!AAA!de passe")
		if got != nil {
			t.Errorf("%v", got)
		}
	})
	t.Run("validate password 1", func(t *testing.T) {
		var user User
		got := user.ValidatePassword("")
		if got == nil {
			t.Errorf("%v", got)
		}
	})
	t.Run("validate password 1", func(t *testing.T) {
		var user User
		got := user.ValidatePassword("1211")
		if got == nil {
			t.Errorf("%v", got)
		}
	})
}
