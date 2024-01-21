package handler

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags 				Api List
// @Security 			JWTAuth
// @Summary 			Create list
// @Description 		Creates a new shopping list and returns its ID
// @Accept  			json
// @Produce  			json
// @Param 				ListInfo			body		model.ShoppingList	true	"Title and description of list"
// @Success 			200 				{object} 	okResponse			 		"Object id"
// @Failure				401					{object}	errorResponse				"Unauthorized"
// @Failure				400					{object}	errorResponse				"Invalid params"
// @Failure				500					{object}	errorResponse				"Failed to create a list"
// @Router 				/api/lists 			[post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.ShoppingList
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "check body: "+err.Error())
		return
	}

	id, err := h.services.List.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, okResponse{Id: id})
}

// @Tags 				Api List
// @Security 			JWTAuth
// @Summary 			Get all lists
// @Description 		Returns a collection of shopping lists
// @Accept				json
// @Produce  			json
// @Success 			200 				{array} 	model.ShoppingList		"Collection of shopping lists"
// @Failure				401					{object}	errorResponse			"Unauthorized"
// @Failure				500					{object}	errorResponse			"Failed to get lists"
// @Router 				/api/lists 			[get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.List.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

// @Tags 				Api List
// @Security 			JWTAuth
// @Summary 			Get list by id
// @Description 		Returns shopping list by ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true				"ID of list"
// @Success 			200 				{object} 	model.ShoppingList		"Shopping list object"
// @Failure				401					{object}	errorResponse			"Unauthorized"
// @Failure				400					{object}	errorResponse			"Invalid params"
// @Failure				500					{object}	errorResponse			"Failed to get the list"
// @Router 				/api/lists/{ID} 	[get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.List.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Tags 				Api List
// @Security 			JWTAuth
// @Summary 			Update list by id
// @Description 		Updates shopping list by ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true						"ID of list"
// @Param 				ListInfo			body		model.UpdateListInput	true	"Title or/and description of list"
// @Success 			200 				{object} 	okResponse						"Object id"
// @Failure				401					{object}	errorResponse					"Unauthorized"
// @Failure				400					{object}	errorResponse					"Invalid params"
// @Failure				500					{object}	errorResponse					"Failed to update the list"
// @Router 				/api/lists/{ID} 	[put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "check body: "+err.Error())
		return
	}

	err = h.services.List.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, okResponse{Id: id})
}

// @Tags 				Api List
// @Security 			JWTAuth
// @Summary 			Delete list by id
// @Description 		Deletes shopping list by ID
// @Accept  			json
// @Produce  			json
// @Param 				ID					path		int	true				"ID of list"
// @Success 			200 				{object} 	okResponse				"Object id"
// @Failure				401					{object}	errorResponse			"Unauthorized"
// @Failure				400					{object}	errorResponse			"Invalid params"
// @Failure				500					{object}	errorResponse			"Failed to delete the list"
// @Router 				/api/lists/{ID} 	[delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.List.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, okResponse{Id: id})
}
