package router

import (
	todo "github.com/dhope-nagesh/gin_todo/pkg/service/v1"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// Register all endpoints here
	v1 := router.Group("v1")
	{
		v1.POST("/todos", todo.AddTodo)
		v1.GET("/todos", todo.ListTodo)
	}
	return router
}
