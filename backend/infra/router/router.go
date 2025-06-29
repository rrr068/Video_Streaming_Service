package router

import (
	"github.com/gin-gonic/gin"
	"github.com/example/gin-backend/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/users", handler.GetUsers)
		// api.POST("/users", handler.CreateUser)
	}

	return r
}