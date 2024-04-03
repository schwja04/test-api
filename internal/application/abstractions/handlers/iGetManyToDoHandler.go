package handlers

import "github.com/schwja04/test-api/internal/domain"

type IGetManyToDoHandler interface {
	Handle() ([]domain.ToDo, error)
}
