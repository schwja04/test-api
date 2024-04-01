package repositories

import (
	"context"

	"github.com/google/uuid"

	repoAbstractions "github.com/schwja04/test-api/application/abstractions/repositories"
	domain_models "github.com/schwja04/test-api/domain"
	"github.com/schwja04/test-api/infrastructure/factories"
)

type ToDoPostgresRepository struct {
	connectionFactory factories.IConnectionFactory
}

func NewToDoPostgresRepository(connectionFactory factories.IConnectionFactory) repoAbstractions.IToDoRepository {
	return &ToDoPostgresRepository{connectionFactory: connectionFactory}
}

func (r *ToDoPostgresRepository) Create(todo domain_models.ToDo) (uuid.UUID, error) {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return uuid.Nil, err
	}
	defer dbConnection.Release()

	_, err = dbConnection.Exec(
		context.Background(),
		"INSERT INTO todos (id, title, content, created_at, updated_at, assignee_id) VALUES ($1, $2, $3, $4, $5, $6)",
		todo.Id, todo.Title, todo.Content, todo.CreatedAt, todo.UpdatedAt, todo.AssigneeId)

	if err != nil {
		return uuid.Nil, err
	}

	return todo.Id, nil
}

func (r *ToDoPostgresRepository) Delete(id uuid.UUID) error {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return err
	}
	defer dbConnection.Release()

	_, err = dbConnection.Exec(context.Background(), "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ToDoPostgresRepository) Get(id uuid.UUID) (domain_models.ToDo, error) {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return domain_models.ToDo{}, err
	}
	defer dbConnection.Release()

	var todo domain_models.ToDo
	err = dbConnection.
		QueryRow(context.Background(), "SELECT id, title, content, created_at, updated_at, assignee_id FROM todos WHERE id = $1", id).
		Scan(&todo.Id, &todo.Title, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt, &todo.AssigneeId)

	if err != nil {
		return domain_models.ToDo{}, err
	}

	return todo, nil
}

func (r *ToDoPostgresRepository) GetAll() ([]domain_models.ToDo, error) {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return nil, err
	}
	defer dbConnection.Release()

	rows, err := dbConnection.Query(context.Background(), "SELECT id, title, content, created_at, updated_at, assignee_id FROM todos")
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

func (r *ToDoPostgresRepository) Update(todo domain_models.ToDo) error {
	dbConnection, err := r.connectionFactory.GetConnection()
	if err != nil {
		return err
	}
	defer dbConnection.Release()

	_, err = dbConnection.Exec(
		context.Background(),
		"UPDATE todos SET title = $1, content = $2, updated_at = $3, assignee_id = $4 WHERE id = $5",
		todo.Title, todo.Content, todo.UpdatedAt, todo.AssigneeId, todo.Id)

	if err != nil {
		return err
	}

	return nil
}
