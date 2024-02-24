package router

import (
	"github.com/Kiluameta/minha-feira-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", handler.ListUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
		v1.DELETE("/user", handler.DeleteUserHandler)
		v1.PUT("/user", handler.UpdateUserHandler)
	}
}
