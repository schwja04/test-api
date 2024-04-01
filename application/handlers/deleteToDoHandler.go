package handlers

import (
	"github.com/google/uuid"
	"github.com/schwja04/test-api/application/abstractions/repositories"
)

type DeleteToDoHandler struct {
	toDoRepository repositories.IToDoRepository
}

func NewDeleteToDoHandler(toDoRepository repositories.IToDoRepository) DeleteToDoHandler {
	return DeleteToDoHandler{toDoRepository: toDoRepository}
}

func (h DeleteToDoHandler) Handle(id uuid.UUID) error {
	err := h.toDoRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
