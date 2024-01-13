package service

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/elecshen/shopping_list/internal/repository"
)

type ShoppingListService struct {
	repo repository.List
}

func NewShoppingListService(repo repository.List) *ShoppingListService {
	return &ShoppingListService{repo: repo}
}

func (s *ShoppingListService) Create(userId int, list model.ShoppingList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ShoppingListService) GetAll(userId int) ([]model.ShoppingList, error) {
	return s.repo.GetAll(userId)
}

func (s *ShoppingListService) GetById(userId, listId int) (model.ShoppingList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ShoppingListService) Update(userId, listId int, input model.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

func (s *ShoppingListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
