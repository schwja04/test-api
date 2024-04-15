package handlers

import (
	"context"

	"github.com/schwja04/test-api/internal/application/abstractions/handlers"
	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
	"github.com/schwja04/test-api/internal/domain"
)

type GetManyToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewGetManyToDoHandler(toDoRepository repositories.IToDoRepository) handlers.IGetManyToDoHandler {
	return &GetManyToDoHandler{toDoRepository: toDoRepository}
}

func (h *GetManyToDoHandler) Handle(ctx context.Context) ([]domain.ToDo, error) {
	todos, err := h.toDoRepository.GetAll(ctx)

	if err != nil {
		return []domain.ToDo{}, err
	}

	return todos, nil
}
