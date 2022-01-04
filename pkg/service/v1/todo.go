package todo

import (
	"github.com/dhope-nagesh/gin_todo/pkg/service/tododb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTodo(c *gin.Context) {
	var todo tododb.Todo
	c.BindJSON(&todo)
	t := tododb.TodoListInstace.Add(todo)
	c.JSON(http.StatusOK, t)
}

func ListTodo(c *gin.Context) {
	todos := tododb.TodoListInstace.List()
	c.JSON(http.StatusOK, todos)
}

