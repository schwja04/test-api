package domain

import (
	"time"

	"github.com/google/uuid"
)

type ToDo struct {
	Id         uuid.UUID
	Title      string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	AssigneeId string
}
