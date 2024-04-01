package handlers

import (
	"time"

	"github.com/google/uuid"
	"github.com/schwja04/test-api/application/abstractions/repositories"
	"github.com/schwja04/test-api/application/commands"
	"github.com/schwja04/test-api/domain"
)

type AddToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewAddToDoHandler(toDoRepository repositories.IToDoRepository) AddToDoHandler {
	return AddToDoHandler{toDoRepository: toDoRepository}
}

func (h AddToDoHandler) Handle(command commands.AddToDoCommand) (uuid.UUID, error) {
	currentTime := time.Now()

	toDo := domain.ToDo{
		Id:         uuid.New(),
		Title:      command.Title,
		Content:    command.Content,
		AssigneeId: command.AssigneeId,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	id, err := h.toDoRepository.Create(toDo)

	return id, err
}
