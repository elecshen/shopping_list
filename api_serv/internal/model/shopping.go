package model

type ShoppingList struct {
	Id          int    `json:"id" db:"ID"`
	Title       string `json:"title" db:"Title" binding:"required"`
	Description string `json:"description" db:"Description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ShoppingItem struct {
	Id          int    `json:"id" db:"ID"`
	Title       string `json:"title" db:"Title"`
	Description string `json:"description" db:"Description"`
	Checked     bool   `json:"checked" db:"Checked"`
}

type ListsItem struct {
	Id      int
	ListsId int
	ItemId  int
}
