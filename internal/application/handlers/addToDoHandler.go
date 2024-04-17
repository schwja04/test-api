package handlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
	"github.com/schwja04/test-api/internal/application/commands"
	"github.com/schwja04/test-api/internal/domain"
)

type AddToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewAddToDoHandler(toDoRepository repositories.IToDoRepository) *AddToDoHandler {
	return &AddToDoHandler{toDoRepository: toDoRepository}
}

func (h *AddToDoHandler) Handle(ctx context.Context, command commands.AddToDoCommand) (uuid.UUID, error) {
	currentTime := time.Now()

	toDo := domain.ToDo{
		Id:         uuid.New(),
		Title:      command.Title,
		Content:    command.Content,
		AssigneeId: command.AssigneeId,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}

	id, err := h.toDoRepository.Create(ctx, toDo)

	return id, err
}
