package repository

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string) (model.User, error)
}

type List interface {
	Create(userId int, list model.ShoppingList) (int, error)
	GetAll(userId int) ([]model.ShoppingList, error)
	GetById(userId, listId int) (model.ShoppingList, error)
	Update(userId, listId int, input model.UpdateListInput) error
	Delete(userId, listId int) error
}

type Item interface {
	Create(listId int, item model.ShoppingItem) (int, error)
	GetAll(userId, listId int) ([]model.ShoppingItem, error)
	GetById(userId, itemId int) (model.ShoppingItem, error)
	Update(userId, itemId int, input model.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewShoppingListPostgres(db),
		Item:          NewShoppingItemPostgres(db),
	}
}
