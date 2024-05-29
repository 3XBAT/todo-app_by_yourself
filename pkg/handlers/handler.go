package handlers

import (
	_ "net/http"

	"github.com/3XBAT/todo-app_by_yourself/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services} 
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentify)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/lists", h.getAllLists)
			lists.GET("/lists/:id", h.getListById)
			lists.POST("/lists", h.createList)
			lists.PUT("lists/:id", h.updateList)
			lists.DELETE("lists/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)

				
			}
		}

		items := api.Group("items")
		{
			items.PUT("/:id", h.updateItem) 
			items.GET("/:id", h.getItemById)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
