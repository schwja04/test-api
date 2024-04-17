package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
	"github.com/schwja04/test-api/internal/domain"
)

type GetSingleToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewGetSingleToDoHandler(toDoRepository repositories.IToDoRepository) *GetSingleToDoHandler {
	return &GetSingleToDoHandler{toDoRepository: toDoRepository}
}

func (h *GetSingleToDoHandler) Handle(ctx context.Context, id uuid.UUID) (domain.ToDo, error) {
	todo, err := h.toDoRepository.Get(ctx, id)

	if err != nil {
		return domain.ToDo{}, err
	}

	return todo, nil
}
