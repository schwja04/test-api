package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/application/abstractions/handlers"
	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
)

type DeleteToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewDeleteToDoHandler(toDoRepository repositories.IToDoRepository) handlers.IDeleteToDoHandler {
	return &DeleteToDoHandler{toDoRepository: toDoRepository}
}

func (h *DeleteToDoHandler) Handle(ctx context.Context, id uuid.UUID) error {
	err := h.toDoRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
