package handlers

import (
	"context"
	"net/http"

	"github.com/3XBAT/todo-app_by_yourself"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	resp, err := h.authClient.Register(context.Background(), input.Name, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//id, err := h.service.Authorization.CreateUser(input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": resp.UserId,
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

	//token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)

	//if err != nil {
	//	logrus.Errorf("Error while generating token:%s", err.Error())
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}

	resp, err := h.authClient.Login(context.Background(), input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": resp.Token,
	})

}
