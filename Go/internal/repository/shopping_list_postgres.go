package repository

import (
	"fmt"
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ShoppingListPostgres struct {
	db *sqlx.DB
}

func NewShoppingListPostgres(db *sqlx.DB) *ShoppingListPostgres {
	return &ShoppingListPostgres{db: db}
}

func (r *ShoppingListPostgres) Create(userId int, list model.ShoppingList) (int, error) {
	var id int

	createListQuery := fmt.Sprintf(`INSERT INTO "Shopping_List" ("Title", "User_id","Description") VALUES ($1, $2, $3) RETURNING "ID"`)
	row := r.db.QueryRow(createListQuery, list.Title, userId, list.Description)
	err := row.Scan(&id)

	return id, err
}

func (r *ShoppingListPostgres) GetAll(userId int) ([]model.ShoppingList, error) {
	var lists []model.ShoppingList

	query := fmt.Sprintf(`SELECT L."ID", L."Title", L."Description" FROM "Shopping_List" L WHERE L."User_id" = $1`)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *ShoppingListPostgres) GetById(userId, listId int) (model.ShoppingList, error) {
	var list model.ShoppingList

	query := fmt.Sprintf(`SELECT L."ID", L."Title", L."Description" FROM "Shopping_List" L WHERE L."User_id" = $1 AND L."ID"=$2`)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *ShoppingListPostgres) Update(userId, listId int, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf(`"Title"=$%d`, argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf(`"Description"=$%d`, argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE "Shopping_List" L SET %s WHERE L."User_id" = $%d AND L."ID" = $%d`,
		setQuery, argId, argId+1)
	args = append(args, userId, listId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ShoppingListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(`DELETE FROM "Shopping_List" L WHERE L."User_id" = $1 AND L."ID" = $2`)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
