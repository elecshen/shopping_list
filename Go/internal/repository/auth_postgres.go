package repository

import (
	"fmt"
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO "User" ("Username", "Salt", "Hash") VALUES ($1, $2, $3) RETURNING "ID"`)
	row := r.db.QueryRow(query, user.Username, user.Salt, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`SELECT "ID", "Salt", "Hash" FROM "User" WHERE "Username" = $1`)
	err := r.db.Get(&user, query, username)

	return user, err
}
