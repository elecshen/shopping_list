package repository

import (
	"fmt"
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ShoppingItemPostgres struct {
	db *sqlx.DB
}

func NewShoppingItemPostgres(db *sqlx.DB) *ShoppingItemPostgres {
	return &ShoppingItemPostgres{db: db}
}

func (r *ShoppingItemPostgres) Create(listId int, item model.ShoppingItem) (int, error) {
	var itemId int

	createItemQuery := fmt.Sprintf(`INSERT INTO "Shopping_Item" ("Title", "List_id", "Description") VALUES ($1, $2, $3) RETURNING "ID"`)
	row := r.db.QueryRow(createItemQuery, item.Title, listId, item.Description)
	err := row.Scan(&itemId)

	return itemId, err
}

func (r *ShoppingItemPostgres) GetAll(userId, listId int) ([]model.ShoppingItem, error) {
	var items []model.ShoppingItem
	query := fmt.Sprintf(`SELECT I."ID", I."Title", I."Description", I."Checked" FROM "Shopping_Item" I INNER JOIN "Shopping_List" L on I."List_id" = L."ID"
									WHERE I."List_id" = $1 AND L."User_id" = $2`)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ShoppingItemPostgres) GetById(userId, itemId int) (model.ShoppingItem, error) {
	var item model.ShoppingItem
	query := fmt.Sprintf(`SELECT I."ID", I."Title", I."Description", I."Checked" FROM "Shopping_Item" I INNER JOIN "Shopping_List" L on I."List_id" = L."ID"
									WHERE I."ID" = $1 AND L."User_id" = $2`)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *ShoppingItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM "Shopping_Item" I USING "Shopping_List" L
									WHERE L."ID" = I."List_id" AND L."User_id" = $1 AND I."ID" = $2`)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *ShoppingItemPostgres) Update(userId, itemId int, input model.UpdateItemInput) error {
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

	if input.Checked != nil {
		setValues = append(setValues, fmt.Sprintf(`"Checked"=$%d`, argId))
		args = append(args, *input.Checked)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE "Shopping_Item" I SET %s FROM "Shopping_List" L
									WHERE L."ID" = I."List_id" AND L."User_id" = $%d AND I."ID" = $%d`,
		setQuery, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
