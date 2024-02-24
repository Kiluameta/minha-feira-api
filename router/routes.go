package router

import (
	"github.com/Kiluameta/minha-feira-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", handler.CreateUserHandler)
	}
}
