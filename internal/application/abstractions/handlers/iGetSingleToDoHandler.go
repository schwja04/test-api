package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/domain"
)

type IGetSingleToDoHandler interface {
	Handle(ctx context.Context, id uuid.UUID) (domain.ToDo, error)
}
