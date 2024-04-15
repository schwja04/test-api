package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"

	repoAbstractions "github.com/schwja04/test-api/internal/application/abstractions/repositories"
	domain_models "github.com/schwja04/test-api/internal/domain"
)

type ToDoInMemoryRepository struct {
	todos map[uuid.UUID]*domain_models.ToDo
}

func NewToDoInMemoryRepository() repoAbstractions.IToDoRepository {
	return &ToDoInMemoryRepository{
		todos: make(map[uuid.UUID]*domain_models.ToDo),
	}
}

func (r *ToDoInMemoryRepository) Create(ctx context.Context, todo domain_models.ToDo) (uuid.UUID, error) {
	return r.upsert(todo)
}

func (r *ToDoInMemoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, ok := r.todos[id]

	if !ok {
		return errors.New("ToDo not found")
	}

	delete(r.todos, id)

	_, ok = r.todos[id]

	if ok {
		return errors.New("ToDo failed to delete")
	}

	return nil
}

func (r *ToDoInMemoryRepository) Get(ctx context.Context, id uuid.UUID) (domain_models.ToDo, error) {
	todo, ok := r.todos[id]

	if !ok {
		return domain_models.ToDo{}, errors.New("ToDo not found")
	}

	return *todo, nil
}

func (r *ToDoInMemoryRepository) GetAll(ctx context.Context) ([]domain_models.ToDo, error) {
	// convert map to slice
	todos := make([]domain_models.ToDo, 0, len(r.todos))

	for _, v := range r.todos {
		todos = append(todos, *v)
	}

	return todos, nil
}

func (r *ToDoInMemoryRepository) Update(ctx context.Context, todo domain_models.ToDo) error {
	_, err := r.upsert(todo)

	return err
}

func (r *ToDoInMemoryRepository) upsert(todo domain_models.ToDo) (uuid.UUID, error) {
	r.todos[todo.Id] = &todo

	return todo.Id, nil
}
