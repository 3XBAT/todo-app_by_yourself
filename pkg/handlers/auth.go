package handlers

import (
	"fmt"
	"net/http"

	"github.com/3XBAT/todo-app_by_yourself"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	fmt.Printf("ID-%d Username-%s, Name-%s, Password-%s\n", input.Id, input.Username, input.Name, input.Password)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	fmt.Printf("ID-%d Username-%s, Name-%s, Password-%s\n",input.Id, input.Username, input.Name, input.Password)


	id, err := h.service.Authorization.CreateUser(input)

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
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("error while binding input(sign-in)")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		logrus.Errorf("Error while generating token:%s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
