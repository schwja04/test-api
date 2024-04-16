package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/schwja04/test-api/internal/api/controllers"
	"github.com/schwja04/test-api/internal/api/routes"
)

func RegisterToDoRoutes(router *gin.Engine, controller *controllers.ToDoController) {
	// Create a new router group
	todos := router.Group(routes.ToDos)

	// Register the routes
	todos.GET("", controller.GetMany)
	todos.POST("", controller.Add)
	todos.DELETE("/:id", controller.Delete)
	todos.GET("/:id", controller.GetSingle)
	todos.PUT("/:id", controller.Update)
}
