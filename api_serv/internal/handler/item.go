package handler

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags 				Api Item
// @Security 			JWTAuth
// @Summary 			Create item
// @Description 		Creates a new item of shopping list and returns its ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true					"ID of list"
// @Param 				ItemInfo			body		model.ShoppingItem	true	"Title and description of item"
// @Success 			200 				{object} 	okResponse			 		"Object id"
// @Failure				401					{object}	errorResponse				"Unauthorized"
// @Failure				400					{object}	errorResponse				"Invalid params"
// @Failure				500					{object}	errorResponse				"Failed to create item"
// @Router 				/api/lists/{ID}/items 			[post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input model.ShoppingItem
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Item.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, okResponse{Id: id})
}

// @Tags 				Api Item
// @Security 			JWTAuth
// @Summary 			Get all items of the list
// @Description 		Returns a collection of items
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true					"ID of list"
// @Success 			200 				{array} 	model.ShoppingItem			"Collection of items"
// @Failure				401					{object}	errorResponse				"Unauthorized"
// @Failure				400					{object}	errorResponse				"Invalid params"
// @Failure				500					{object}	errorResponse				"Failed to get items"
// @Router 				/api/lists/{ID}/items 			[get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.Item.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Tags 				Api Item
// @Security 			JWTAuth
// @Summary 			Get item by id
// @Description 		Returns shopping item by ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true						"ID of item"
// @Success 			200 				{object} 	okResponse						"Object id"
// @Failure				401					{object}	errorResponse					"Unauthorized"
// @Failure				400					{object}	errorResponse					"Invalid params"
// @Failure				500					{object}	errorResponse					"Failed to get the item"
// @Router 				/api/items/{ID} 	[get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.Item.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Tags 				Api Item
// @Security 			JWTAuth
// @Summary 			Update item by id
// @Description 		Updates shopping item by ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true						"ID of item"
// @Param 				ListInfo			body		model.UpdateItemInput	true	"Title or/and description of item"
// @Success 			200 				{object} 	okResponse						"Object id"
// @Failure				401					{object}	errorResponse					"Unauthorized"
// @Failure				400					{object}	errorResponse					"Invalid params"
// @Failure				500					{object}	errorResponse					"Failed to update the item"
// @Router 				/api/items/{ID} 	[put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input model.UpdateItemInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Item.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, okResponse{Id: id})
}

// @Tags 				Api Item
// @Security 			JWTAuth
// @Summary 			Delete item by id
// @Description 		Deletes shopping item by ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true						"ID of item"
// @Success 			200 				{object} 	okResponse						"Object id"
// @Failure				401					{object}	errorResponse					"Unauthorized"
// @Failure				400					{object}	errorResponse					"Invalid params"
// @Failure				500					{object}	errorResponse					"Failed to delete the item"
// @Router 				/api/items/{ID} 	[delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.services.Item.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, okResponse{Id: id})
}
