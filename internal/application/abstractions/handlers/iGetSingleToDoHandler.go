package handlers

import (
	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/domain"
)

type IGetSingleToDoHandler interface {
	Handle(id uuid.UUID) (domain.ToDo, error)
}
