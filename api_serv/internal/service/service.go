package service

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/elecshen/shopping_list/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(userId int, list model.ShoppingList) (int, error)
	GetAll(userId int) ([]model.ShoppingList, error)
	GetById(userId, listId int) (model.ShoppingList, error)
	Update(userId, listId int, input model.UpdateListInput) error
	Delete(userId, listId int) error
}

type Item interface {
	Create(userId, listId int, item model.ShoppingItem) (int, error)
	GetAll(userId, listId int) ([]model.ShoppingItem, error)
	GetById(userId, itemId int) (model.ShoppingItem, error)
	Update(userId, itemId int, input model.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		List:          NewShoppingListService(repos.List),
		Item:          NewShoppingItemService(repos.Item, repos.List),
	}
}
