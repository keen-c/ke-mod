package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/keen-c/modular/shared/database"
)
var UserStore = NewUserStorer(database.InitDB)

func PostCreate(w http.ResponseWriter, r *http.Request) {
	user := User{
		email:    strings.ToLower(r.FormValue("email")),
		password: r.FormValue("password"),
	}
	repeat := r.FormValue("repeat")
	m := user.Validate(repeat)
	if m != nil {
		fmt.Fprintf(w, "%v", m)
		return
	}
	err := user.ValidatePassword(user.password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	id, err := UserStore.Create(r.Context(), user.email, user.password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
