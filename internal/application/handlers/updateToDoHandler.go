package handlers

import (
	"context"
	"time"

	"github.com/schwja04/test-api/internal/application/abstractions/handlers"
	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
	"github.com/schwja04/test-api/internal/application/commands"
)

type UpdateToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewUpdateToDoHandler(toDoRepository repositories.IToDoRepository) handlers.IUpdateToDoHandler {
	return &UpdateToDoHandler{toDoRepository: toDoRepository}
}

func (h *UpdateToDoHandler) Handle(ctx context.Context, command commands.UpdateToDoCommand) error {
	todo, err := h.toDoRepository.Get(ctx, command.Id)

	if err != nil {
		return err
	}

	todo.Title = command.Title
	todo.Content = command.Content
	todo.AssigneeId = command.AssigneeId
	todo.UpdatedAt = time.Now()

	err = h.toDoRepository.Update(ctx, todo)

	if err != nil {
		return err
	}

	return nil
}
