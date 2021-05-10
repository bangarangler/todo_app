package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bangarangler/todo_app/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func goDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env")
	}
	return os.Getenv(key)
}

var (
	port = goDotEnvVar("PORT")
	r    = repository.NewRepository()
)

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": r.ListTodo()})
}

func AddTodo(c *gin.Context) {
	var input repository.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r.AddTodo(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func main() {
	if port == "" {
		port = strconv.Itoa(8080)
	}
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/todos", GetTodos)
		v1.POST("/todos", AddTodo)
	}
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
