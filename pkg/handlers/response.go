package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Messgae string `json:"Message"`
}

type statusReasponse struct {//структура для ответа на эндпоинт удаления списка
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(status, errorResponse{message})
}
