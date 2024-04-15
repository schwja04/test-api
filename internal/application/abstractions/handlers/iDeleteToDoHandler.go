package handlers

import (
	"context"

	"github.com/google/uuid"
)

type IDeleteToDoHandler interface {
	Handle(ctx context.Context, id uuid.UUID) error
}
