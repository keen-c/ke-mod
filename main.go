package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/keen-c/modular/shared/database"
	"github.com/keen-c/modular/user"
)
func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic("cant load env variables :", err)
		return
	}
	db , err := database.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer func(){
		defer db.Close()
	}()
	apiUser := user.NewUserHandlerAPi(user.NewUserStorer(db))

	r := chi.NewRouter()
	r.Get("/create", apiUser.GetCreate)
	r.Post("/create", apiUser.PostCreate)
	fmt.Println("the app is running : http://localhost:9292/create")
	if err := http.ListenAndServe(":9292", r); err != nil {
		panic(err)
	}
}