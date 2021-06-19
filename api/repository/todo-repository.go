package repository

import (
	"context"
	"errors"
	"go-vue-todo/model"
)

var ErrEntityNotFound = errors.New("entity not found")

type TodoRepository interface {
	List(context.Context, string) ([]*model.Todo, error)
	GetById(context.Context, int) (*model.Todo, error)
	Create(context.Context, *model.Todo) error
	Update(context.Context, *model.Todo) error
	Delete(context.Context, int) error
}