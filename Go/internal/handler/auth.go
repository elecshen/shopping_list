package handler

import (
	"github.com/elecshen/shopping_list/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags 				Authorization
// @Summary 			Registration
// @Description 		Creates a user account and returns its ID
// @Accept  			json
// @Produce  			json
// @Param 				User				body		model.User	true	"Account info"
// @Success 			200 				{object} 	okResponse 			"Object id"
// @Failure				400					{object}	errorResponse		"Invalid body content"
// @Failure				500					{object}	errorResponse		"Failed to create a user"
// @Router 				/auth/sign-up 		[post]
func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, okResponse{Id: id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Tags 				Authorization
// @Summary 			Login
// @Description 		Authorizes the user and returns the access token
// @Accept  			json
// @Produce  			json
// @Param 				SignInInput			body		signInInput	true	"Username and password"
// @Success 			200 				{object} 	tokenResponse 		"Access token"
// @Failure				400					{object}	errorResponse		"Invalid body content"
// @Failure				500					{object}	errorResponse		"Failed to generate token"
// @Router 				/auth/sign-in 		[post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{Token: token})
}
