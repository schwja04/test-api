package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/schwja04/test-api/internal/api"
	"github.com/schwja04/test-api/internal/api/controllers"
	"github.com/schwja04/test-api/internal/application/handlers"
	"github.com/schwja04/test-api/internal/infrastructure/factories"
	"github.com/schwja04/test-api/internal/infrastructure/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Repository Dependencies
	connectionFactory := factories.NewPostgresConnectionFactory("postgres://postgres:postgres@my-postgres:5432/todo_db")
	defer connectionFactory.Close()

	// Repositories
	todoRepo := repositories.NewToDoPostgresRepository(connectionFactory)

	// Application Handlers
	addHandler := handlers.NewAddToDoHandler(todoRepo)
	deleteHandler := handlers.NewDeleteToDoHandler(todoRepo)
	getSingleHandler := handlers.NewGetSingleToDoHandler(todoRepo)
	getManyHandler := handlers.NewGetManyToDoHandler(todoRepo)
	updateHandler := handlers.NewUpdateToDoHandler(todoRepo)

	// API Controllers
	todoController := controllers.NewToDoController(
		addHandler, deleteHandler, getManyHandler, getSingleHandler, updateHandler)

	// API
	router := gin.Default()

	// Routes
	RegisterToDoRoutes(router, todoController)

	apiPort := os.Getenv("API_PORT")
	router.Run(":" + apiPort)
}

func RegisterToDoRoutes(router *gin.Engine, controller *controllers.ToDoController) {
	router.GET(routes.ToDos, controller.GetMany)
	router.POST(routes.ToDos, controller.Add)
	router.DELETE(routes.ToDoById, controller.Delete)
	router.GET(routes.ToDoById, controller.GetSingle)
	router.PUT(routes.ToDoById, controller.Update)
}
