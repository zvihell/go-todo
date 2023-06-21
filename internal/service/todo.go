package service

import (
	"context"
	"todo-api/internal/models"
)

type TodoRepository interface {
	Create(ctx context.Context, todo models.Todo) error
	GetAll(ctx context.Context) ([]models.Todo, error)
	GetByID(ctx context.Context, id int) (models.Todo, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, todo models.Todo) error
}
type Todo struct {
	repo TodoRepository
}

func NewTodo(repo TodoRepository) *Todo {
	return &Todo{
		repo: repo,
	}
}

func (t *Todo) Create(ctx context.Context, todo models.Todo) error {
	return t.repo.Create(ctx, todo)
}

func (t *Todo) GetAll(ctx context.Context) ([]models.Todo, error) {
	return t.repo.GetAll(ctx)
}

func (t *Todo) GetByID(ctx context.Context, id int) (models.Todo, error) {
	return t.repo.GetByID(ctx, id)
}
func (t *Todo) Delete(ctx context.Context, id int) error {
	return t.repo.Delete(ctx, id)
}

func (t *Todo) Update(ctx context.Context, id int, todo models.Todo) error {
	return t.repo.Update(ctx, id, todo)
}
