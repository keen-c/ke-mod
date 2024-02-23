package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/keen-c/modular/user"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("cant load env variables :", err)
		return
	}
}
func main() {
	Init()
	var user user.User
	err := user.ValidatePassword("pa!sswSrd")
	fmt.Println(err)
}
