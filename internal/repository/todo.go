package repository

import (
	"context"
	"database/sql"
	"errors"
	"todo-api/internal/models"
)

var (
	ErrTodoNotFound = errors.New("Todo not found")
)

type Todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *Todo {
	return &Todo{db: db}
}

func (t *Todo) Create(ctx context.Context, todo models.Todo) error {
	_, err := t.db.Exec("INSERT INTO todos (description) VALUES($1)", todo.Description)
	return err
}

func (t *Todo) GetAll(ctx context.Context) ([]models.Todo, error) {
	rows, err := t.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Description); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()

}

func (t *Todo) GetByID(ctx context.Context, id int) (models.Todo, error) {
	var todo models.Todo
	err := t.db.QueryRow("SELECT id, description FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Description)
	if err == sql.ErrNoRows {
		return todo, ErrTodoNotFound
	}
	return todo, err
}

func (t *Todo) Delete(ctx context.Context, id int) error {
	_, err := t.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (t *Todo) Update(ctx context.Context, id int, todo models.Todo) error {
	_, err := t.db.Exec("UPDATE todos SET description=$1 WHERE id = $2", todo.Description, id)
	return err
}
