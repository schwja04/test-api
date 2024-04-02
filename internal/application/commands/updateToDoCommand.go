package commands

import (
	"github.com/google/uuid"
)

type UpdateToDoCommand struct {
	Id         uuid.UUID
	Title      string
	Content    string
	AssigneeId string
}
