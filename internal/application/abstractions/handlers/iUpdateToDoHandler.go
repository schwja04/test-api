package handlers

import "github.com/schwja04/test-api/internal/application/commands"

type IUpdateToDoHandler interface {
	Handle(command commands.UpdateToDoCommand) error
}
