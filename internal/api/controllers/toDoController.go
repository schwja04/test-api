package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/schwja04/test-api/internal/api/contracts"
	"github.com/schwja04/test-api/internal/application/abstractions/handlers"
	"github.com/schwja04/test-api/internal/application/commands"
)

// Need our handlers injected, inject as interfaces
type ToDoController struct {
	addHandler       handlers.IAddToDoHandler
	deleteHandler    handlers.IDeleteToDoHandler
	getSingleHandler handlers.IGetSingleToDoHandler
	getManyHandler   handlers.IGetManyToDoHandler
	updateHandler    handlers.IUpdateToDoHandler
}

func NewToDoController(
	addHandler handlers.IAddToDoHandler,
	deleteHandler handlers.IDeleteToDoHandler,
	getManyHandler handlers.IGetManyToDoHandler,
	getSingleHandler handlers.IGetSingleToDoHandler,
	updateHandler handlers.IUpdateToDoHandler) *ToDoController {
	return &ToDoController{
		addHandler:       addHandler,
		deleteHandler:    deleteHandler,
		getManyHandler:   getManyHandler,
		getSingleHandler: getSingleHandler,
		updateHandler:    updateHandler,
	}
}

func (controller *ToDoController) GetMany(c *gin.Context) {
	todos, err := controller.getManyHandler.Handle()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (controller *ToDoController) Add(c *gin.Context) {
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

	id, err := controller.addHandler.Handle(command)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
	c.Header("Location", fmt.Sprintf("%s/%s", c.Request.URL.Path, id.String()))
}

func (controller *ToDoController) GetSingle(c *gin.Context) {
	todoId := c.Param("id")

	todoUUID, err := uuid.Parse(todoId)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := controller.getSingleHandler.Handle(todoUUID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (controller *ToDoController) Update(c *gin.Context) {
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

	err = controller.updateHandler.Handle(command)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
	c.Header("Location", fmt.Sprintf("%s/%s", c.Request.URL.Path, todoId))
}

func (controller *ToDoController) Delete(c *gin.Context) {
	todoId := c.Param("id")

	todoUUID, err := uuid.Parse(todoId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.deleteHandler.Handle(todoUUID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
