package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json: "id"`
	Item      string `json: "item"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "Limpar o quarto", Completed: false},
	{ID: "2", Item: "Academia", Completed: false},
	{ID: "3", Item: "Dormir", Completed: false},
}

func getTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func addTodo(ctx *gin.Context) {
	var newTodo todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
	}

	todo.Completed = !todo.Completed

	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)

	router.Run("localhost:9090")
}
