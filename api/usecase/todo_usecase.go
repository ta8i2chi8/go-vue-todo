package usecase

import (
	"context"
	"time"

	"go-vue-todo/model"
	"go-vue-todo/repository"
)

type TodoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *TodoUseCase) List(ctx context.Context, qn string) ([]*model.Todo, error) {
	return uc.todoRepo.List(ctx, qn)
}

func (uc *TodoUseCase) GetById(ctx context.Context, id int) (*model.Todo, error) {
	return uc.todoRepo.GetById(ctx, id)
}

type CreateParams struct {
	Title     string
	Context   string
	LimitDate time.Time
}

func (uc *TodoUseCase) Create(ctx context.Context, param *CreateParams) error {
	td := &model.Todo{
		Title:     param.Title,
		Context:   param.Context,
		LimitDate: param.LimitDate,
	}

	return uc.todoRepo.Create(ctx, td)
}

type UpdateParams struct {
	ID        int
	Title     string
	Context   string
	LimitDate time.Time
}

func (uc *TodoUseCase) Update(ctx context.Context, param *UpdateParams) error {
	td, err := uc.todoRepo.GetById(ctx, param.ID)
	if err != nil {
		return err
	}
	td.Title = param.Title
	td.Context = param.Context
	td.LimitDate = param.LimitDate
	return uc.todoRepo.Update(ctx, td)
}

func (uc *TodoUseCase) Delete(ctx context.Context, id int) error {
	return uc.todoRepo.Delete(ctx, id)
}
