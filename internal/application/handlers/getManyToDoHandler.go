package handlers

import (
	"github.com/schwja04/test-api/internal/application/abstractions/repositories"
	"github.com/schwja04/test-api/internal/domain"
)

type GetManyToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewGetManyToDoHandler(toDoRepository repositories.IToDoRepository) GetManyToDoHandler {
	return GetManyToDoHandler{toDoRepository: toDoRepository}
}

func (h GetManyToDoHandler) Handle() ([]domain.ToDo, error) {
	todos, err := h.toDoRepository.GetAll()

	if err != nil {
		return []domain.ToDo{}, err
	}

	return todos, nil
}
