package handlers

import (
	"github.com/google/uuid"
	"github.com/schwja04/test-api/internal/application/commands"
)

type IAddToDoHandler interface {
	Handle(command commands.AddToDoCommand) (uuid.UUID, error)
}
