package service

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/elecshen/shopping_list/internal/repository"
)

type ShoppingItemService struct {
	repo     repository.Item
	listRepo repository.List
}

func NewShoppingItemService(repo repository.Item, listRepo repository.List) *ShoppingItemService {
	return &ShoppingItemService{repo: repo, listRepo: listRepo}
}

func (s *ShoppingItemService) Create(userId, listId int, item model.ShoppingItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exist or does not belong to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *ShoppingItemService) GetAll(userId, listId int) ([]model.ShoppingItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ShoppingItemService) GetById(userId, itemId int) (model.ShoppingItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ShoppingItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ShoppingItemService) Update(userId, itemId int, input model.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
