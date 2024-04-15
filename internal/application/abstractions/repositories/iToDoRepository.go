package repositories

import (
	"context"

	"github.com/google/uuid"
	domain_models "github.com/schwja04/test-api/internal/domain"
)

type IToDoRepository interface {
	Create(ctx context.Context, todo domain_models.ToDo) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (domain_models.ToDo, error)
	GetAll(ctx context.Context) ([]domain_models.ToDo, error)
	Update(ctx context.Context, todo domain_models.ToDo) error
}
