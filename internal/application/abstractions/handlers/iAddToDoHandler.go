package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/application/commands"
)

type IAddToDoHandler interface {
	Handle(ctx context.Context, command commands.AddToDoCommand) (uuid.UUID, error)
}
