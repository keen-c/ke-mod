package user

import (
	"context"
	"database/sql"

	"github.com/keen-c/modular/shared/database"
)

type UserStorer struct {
	DB *sql.DB
}
func NewUserStorer(fn database.ConnectDB) *UserStorer {
	db , err := fn.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return &UserStorer{DB: db}
}
func (us UserStorer) Create(ctx context.Context, email, password string) (string, error) {
	query := `insert into users (email, password) values ($1, $2) returing id`
	var id string
	if err := us.DB.QueryRowContext(ctx, query, email, password).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}
func (us UserStorer) Update() error {
	return nil 
}
func (us UserStorer) Delete(ctx context.Context, email string) error {
	query := `delete from users where email = $1`
	if _, err := us.DB.ExecContext(ctx, query, email); err != nil {
		return err
	}
	return nil
}
func (us UserStorer) Connect(ctx context.Context, email, password string) error {
	query := `select * from users where email = $1`
	var user User
	if err := us.DB.QueryRowContext(ctx, query, email, password).Scan(&user); err != nil {
		return err
	}
	return nil
}
