package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/schwja04/test-api/internal/api/contracts"
	"github.com/schwja04/test-api/internal/application/commands"
	"github.com/schwja04/test-api/internal/application/handlers"
	"github.com/schwja04/test-api/internal/infrastructure/factories"
	"github.com/schwja04/test-api/internal/infrastructure/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	connectionFactory := factories.NewPostgresConnectionFactory("postgres://postgres:postgres@my-postgres:5432/todo_db")
	defer connectionFactory.Close()

	todoRepo := repositories.NewToDoPostgresRepository(connectionFactory)
	addHandler := handlers.NewAddToDoHandler(todoRepo)
	deleteHandler := handlers.NewDeleteToDoHandler(todoRepo)
	getSingleHandler := handlers.NewGetSingleToDoHandler(todoRepo)
	getManyHandler := handlers.NewGetManyToDoHandler(todoRepo)
	updateHandler := handlers.NewUpdateToDoHandler(todoRepo)

	router.GET("/todos", func(c *gin.Context) {

		todos, err := getManyHandler.Handle()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	router.POST("/todos", func(c *gin.Context) {
		var request contracts.AddToDoCommand

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		command := commands.AddToDoCommand{
			Title:      request.Title,
			Content:    request.Content,
			AssigneeId: request.AssigneeId,
		}

		id, err := addHandler.Handle(command)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusCreated)
		c.Header("Location", "/todos/"+id.String())
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		todoId := c.Param("id")

		todoUUID, err := uuid.Parse(todoId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = deleteHandler.Handle(todoUUID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusNoContent)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		todoId := c.Param("id")

		todoUUID, err := uuid.Parse(todoId)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		todo, err := getSingleHandler.Handle(todoUUID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, todo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		todoId := c.Param("id")

		var request contracts.UpdateToDoCommand

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		todoUUID, err := uuid.Parse(todoId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		command := commands.UpdateToDoCommand{
			Id:         todoUUID,
			Title:      request.Title,
			Content:    request.Content,
			AssigneeId: request.AssigneeId,
		}

		err = updateHandler.Handle(command)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusNoContent)
		c.Header("Location", "/todos/"+todoId)
	})

	apiPort := os.Getenv("API_PORT")
	router.Run(":" + apiPort)
}
