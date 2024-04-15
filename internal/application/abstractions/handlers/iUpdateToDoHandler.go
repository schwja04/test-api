package handlers

import (
	"context"

	"github.com/schwja04/test-api/internal/application/commands"
)

type IUpdateToDoHandler interface {
	Handle(ctx context.Context, command commands.UpdateToDoCommand) error
}
