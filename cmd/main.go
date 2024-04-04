package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	routes "github.com/schwja04/test-api/internal/api"
	"github.com/schwja04/test-api/internal/api/controllers"
	"github.com/schwja04/test-api/internal/application/handlers"
	"github.com/schwja04/test-api/internal/infrastructure/factories"
	"github.com/schwja04/test-api/internal/infrastructure/repositories"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

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
	RegisterToDoRoutes(router, todoController)

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		log.Fatal("API_PORT is not set")
	}

	srv := &http.Server{
		Addr:    ":" + apiPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func RegisterToDoRoutes(router *gin.Engine, controller *controllers.ToDoController) {
	router.GET(routes.ToDos, controller.GetMany)
	router.POST(routes.ToDos, controller.Add)
	router.DELETE(routes.ToDoById, controller.Delete)
	router.GET(routes.ToDoById, controller.GetSingle)
	router.PUT(routes.ToDoById, controller.Update)
}
