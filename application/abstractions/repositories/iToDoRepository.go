package repositories

import (
	"github.com/google/uuid"
	domain_models "github.com/schwja04/test-api/domain"
)

type IToDoRepository interface {
	Create(todo domain_models.ToDo) (uuid.UUID, error)
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (domain_models.ToDo, error)
	GetAll() ([]domain_models.ToDo, error)
	Update(todo domain_models.ToDo) error
}
