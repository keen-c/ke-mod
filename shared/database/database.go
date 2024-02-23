package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type Database struct {
	Host     string
	Name     string
	Port     int
	User     string
	Password string
}

type ConnectDB interface {
	InitDB() (*sql.DB, error)
}

func (d *Database) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.User, d.Password, d.Host, d.Port, d.Name)
}
func InitDB() (*sql.DB, error) {
	p, err := strconv.Atoi(os.Getenv("SUP_PORT"))
	if err != nil {
		return nil, err
	}
	dbase := Database{
		Host:     os.Getenv("SUP_HOST"),
		Name:     os.Getenv("SUP_NAME"),
		User:     os.Getenv("SUP_USER"),
		Password: os.Getenv("SUP_PASSWORD"),
	}
	dbase.Port = p
	db, err := sql.Open("postgres", dbase.String())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
