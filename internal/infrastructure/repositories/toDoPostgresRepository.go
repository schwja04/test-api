package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"

	repoAbstractions "github.com/schwja04/test-api/internal/application/abstractions/repositories"
	domain_models "github.com/schwja04/test-api/internal/domain"
	"github.com/schwja04/test-api/internal/infrastructure/postgres/factories"
)

type ToDoPostgresRepository struct {
	connectionFactory factories.IConnectionFactory
}

func NewToDoPostgresRepository(connectionFactory factories.IConnectionFactory) repoAbstractions.IToDoRepository {
	return &ToDoPostgresRepository{connectionFactory: connectionFactory}
}

func (r *ToDoPostgresRepository) Create(ctx context.Context, todo domain_models.ToDo) (uuid.UUID, error) {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer dbConnection.Release()

	dbCtx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	_, err = dbConnection.Exec(
		dbCtx,
		"INSERT INTO todos (id, title, content, created_at, updated_at, assignee_id) VALUES ($1, $2, $3, $4, $5, $6)",
		todo.Id, todo.Title, todo.Content, todo.CreatedAt, todo.UpdatedAt, todo.AssigneeId)

	if err != nil {
		return uuid.Nil, err
	}

	return todo.Id, nil
}

func (r *ToDoPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return err
	}
	defer dbConnection.Release()

	dbCtx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	_, err = dbConnection.Exec(dbCtx, "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ToDoPostgresRepository) Get(ctx context.Context, id uuid.UUID) (domain_models.ToDo, error) {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return domain_models.ToDo{}, err
	}
	defer dbConnection.Release()

	dbCtx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	var todo domain_models.ToDo
	err = dbConnection.
		QueryRow(dbCtx, "SELECT id, title, content, created_at, updated_at, assignee_id FROM todos WHERE id = $1", id).
		Scan(&todo.Id, &todo.Title, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt, &todo.AssigneeId)

	if err != nil {
		return domain_models.ToDo{}, err
	}

	return todo, nil
}

func (r *ToDoPostgresRepository) GetAll(ctx context.Context) ([]domain_models.ToDo, error) {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return nil, err
	}
	defer dbConnection.Release()

	dbCtx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	rows, err := dbConnection.Query(dbCtx, "SELECT id, title, content, created_at, updated_at, assignee_id FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]domain_models.ToDo, 0, 10)
	for rows.Next() {
		var todo domain_models.ToDo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt, &todo.AssigneeId)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *ToDoPostgresRepository) Update(ctx context.Context, todo domain_models.ToDo) error {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return err
	}
	defer dbConnection.Release()

	dbCtx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	_, err = dbConnection.Exec(
		dbCtx,
		"UPDATE todos SET title = $1, content = $2, updated_at = $3, assignee_id = $4 WHERE id = $5",
		todo.Title, todo.Content, todo.UpdatedAt, todo.AssigneeId, todo.Id)

	if err != nil {
		return err
	}

	return nil
}
