package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/schwja04/test-api/internal/api/controllers"
	"github.com/schwja04/test-api/internal/api/routes"
	"github.com/schwja04/test-api/internal/application/handlers"
	"github.com/schwja04/test-api/internal/infrastructure/postgres/builders"
	"github.com/schwja04/test-api/internal/infrastructure/postgres/factories"
	"github.com/schwja04/test-api/internal/infrastructure/repositories"
	"github.com/schwja04/test-api/packages/otel"
	ginOtelMiddleware "github.com/schwja04/test-api/packages/otel/gin"
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

	providerShutdown, err := otel.InitTraceProvider(ctx, "todo-api", false)
	if err != nil {
		log.Warning(fmt.Errorf("failed initializing otel trace provider: %w", err))
	}
	defer func() {
		if err := providerShutdown(ctx); err != nil {
			log.Warning(fmt.Errorf("failed shutting down otel trace provider: %w", err))
		}
	}()

	// I don't like this implementation, but default values should not be set in the builder
	// I could push it as a string and parse it in the builder, but it would be unexpected that the builder
	// would throw an error if the value is not a valid integer, AT LEAST IN THE CURRENT IMPLEMENTATION.
	pgPort := 5432
	if os.Getenv("PG_PORT") != "" {
		parsedPgPort, err := strconv.Atoi(os.Getenv("PG_PORT"))
		if err != nil {
			log.Fatal("PG_PORT is not a valid integer")
		}
		pgPort = parsedPgPort
	}

	connectionString := builders.NewConnectionStringBuilder().
		WithUser(os.Getenv("PG_USER")).
		WithPassword(os.Getenv("PG_PASSWORD")).
		WithHost(os.Getenv("PG_HOST")).
		WithPort(pgPort).
		WithDatabase(os.Getenv("PG_DATABASE")).
		Build()

	// Repository Dependencies
	connectionFactory := factories.NewPostgresConnectionFactory(connectionString)
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
	router.Use(
		gin.Recovery(),
		otelgin.Middleware("todo-api"),
		ginOtelMiddleware.FriendlyOtelMapping(),
	)
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
