package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/keen-c/modular/shared"
	"github.com/keen-c/modular/views/pages"
	
)

type userStore interface {
	Create(ctx context.Context, email, password string) (string, error)
}
type UserHandlerApi struct {
	userStore userStore
}

func NewUserHandlerAPi(userStore userStore) *UserHandlerApi {
	return &UserHandlerApi{userStore: userStore}
}

func (uh *UserHandlerApi) GetCreate(w http.ResponseWriter, r *http.Request) {
	if err := pages.Inscription(ErrorsInscription{}, UserCreate{}).Render(r.Context(), w); err != nil {
		fmt.Println(err)
	}
}

func (uh *UserHandlerApi) PostCreate(w http.ResponseWriter, r *http.Request) {
	user := UserCreate{
		Email:    strings.ToLower(r.FormValue("email")),
		Password: r.FormValue("password"),
	}
	m := user.Validate(r.FormValue("repeat-password"))
	if m != nil {
		pages.Inscription(*m, user).Render(r.Context(), w)
		return
	}
	if err := user.ValidatePassword(user.Password); err != nil {
		fmt.Println("user.Validate")
		pages.Inscription(pages.ErrorsInscription{
			Password: err.Error(),
		}, user).Render(r.Context(), w)
		return
	}
	if err := user.HashPassword(); err != nil {
		shared.WriteError(w)
		return
	}
	id, err := uh.userStore.Create(r.Context(), user.Email, user.Password)
	if err != nil {
		shared.WriteError(w)
		return
	}
	fmt.Println("id: ", id)

}
