package handlers

import (
	"time"

	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
	"github.com/schwja04/test-api/internal/application/commands"
)

type UpdateToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewUpdateToDoHandler(toDoRepository repositories.IToDoRepository) UpdateToDoHandler {
	return UpdateToDoHandler{toDoRepository: toDoRepository}
}

func (h UpdateToDoHandler) Handle(command commands.UpdateToDoCommand) error {
	todo, err := h.toDoRepository.Get(command.Id)

	if err != nil {
		return err
	}

	todo.Title = command.Title
	todo.Content = command.Content
	todo.AssigneeId = command.AssigneeId
	todo.UpdatedAt = time.Now()

	err = h.toDoRepository.Update(todo)

	if err != nil {
		return err
	}

	return nil
}
