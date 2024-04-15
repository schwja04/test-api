package handlers

import (
	"context"

	"github.com/schwja04/test-api/internal/domain"
)

type IGetManyToDoHandler interface {
	Handle(ctx context.Context) ([]domain.ToDo, error)
}
