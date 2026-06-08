package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	handler *Handler,
) {
	users := router.Group("/users")
	{
		users.POST("", handler.Create)
	}
}
