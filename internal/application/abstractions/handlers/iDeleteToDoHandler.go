package handlers

import "github.com/google/uuid"

type IDeleteToDoHandler interface {
	Handle(id uuid.UUID) error
}
