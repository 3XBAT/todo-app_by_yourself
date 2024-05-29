package handlers

import (
	"net/http"

	"github.com/3XBAT/todo-app_by_yourself"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) { 
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	 id, err := h.service.Authorization.CreateUser(input); 

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type singInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"passwrd" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	 token, err := h.service.Authorization.GenerateToken(input.Username, input.Password); 

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

} 
