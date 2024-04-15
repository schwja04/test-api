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

func (controller *ToDoController) GetMany(ctx *gin.Context) {
	todos, err := controller.getManyHandler.Handle(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (controller *ToDoController) Add(ctx *gin.Context) {
	var request contracts.AddToDoCommand

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	command := commands.AddToDoCommand{
		Title:      request.Title,
		Content:    request.Content,
		AssigneeId: request.AssigneeId,
	}

	id, err := controller.addHandler.Handle(ctx, command)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
	ctx.Header("Location", fmt.Sprintf("%s/%s", ctx.Request.URL.Path, id.String()))
}

func (controller *ToDoController) GetSingle(ctx *gin.Context) {
	todoId := ctx.Param("id")

	todoUUID, err := uuid.Parse(todoId)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := controller.getSingleHandler.Handle(ctx, todoUUID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (controller *ToDoController) Update(ctx *gin.Context) {
	todoId := ctx.Param("id")

	var request contracts.UpdateToDoCommand

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoUUID, err := uuid.Parse(todoId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

	err = controller.updateHandler.Handle(ctx, command)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
	ctx.Header("Location", fmt.Sprintf("%s/%s", ctx.Request.URL.Path, todoId))
}

func (controller *ToDoController) Delete(ctx *gin.Context) {
	todoId := ctx.Param("id")

	todoUUID, err := uuid.Parse(todoId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = controller.deleteHandler.Handle(ctx, todoUUID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
