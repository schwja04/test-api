package handlers

import (
	"github.com/google/uuid"
	"github.com/schwja04/test-api/application/abstractions/repositories"
	"github.com/schwja04/test-api/domain"
)

type GetSingleToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewGetSingleToDoHandler(toDoRepository repositories.IToDoRepository) GetSingleToDoHandler {
	return GetSingleToDoHandler{toDoRepository: toDoRepository}
}

func (h GetSingleToDoHandler) Handle(id uuid.UUID) (domain.ToDo, error) {
	todo, err := h.toDoRepository.Get(id)

	if err != nil {
		return domain.ToDo{}, err
	}

	return todo, nil
}
