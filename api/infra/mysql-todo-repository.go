package infra

import (
	"context"
	"fmt"

	"go-vue-todo/model"
	"go-vue-todo/repository"

	"github.com/jmoiron/sqlx"
)

type MySQLTodoRepository struct {
	DB *sqlx.DB
}

func (r *MySQLTodoRepository) List(ctx context.Context, qn string) ([]*model.Todo, error) {
	qn = fmt.Sprintf("%s%%", qn)
	rows, err := r.DB.Queryx("SELECT id, title, context, limitDate, createdAt, updatedAt FROM todos WHERE title LIKE ? ORDER BY id ASC", qn)
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo

	for rows.Next() {
		var todo model.Todo

		err := rows.StructScan(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	return todos, nil
}

func (r *MySQLTodoRepository) GetById(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo

	err := r.DB.QueryRowx("SELECT id, title, context, limitDate, createdAt, updatedAt FROM todos WHERE id=?", id).StructScan(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *MySQLTodoRepository) Create(ctx context.Context, td *model.Todo) error {
	result, err := r.DB.ExecContext(ctx, "INSERT INTO todos(title, context, limitDate) VALUES(?, ?, ?)", td.Title, td.Context, td.LimitDate)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err !=nil {
		return err
	}
	if rows != 1 {
		return repository.ErrEntityNotFound
	}
	return nil
}

func (r *MySQLTodoRepository) Update(ctx context.Context, td *model.Todo) error {
	result, err := r.DB.ExecContext(ctx, "UPDATE todos SET title=?, context=?, limitDate=? WHERE id=?", td.Title, td.Context, td.LimitDate, td.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.ErrEntityNotFound
	}
	return nil
}

func (r *MySQLTodoRepository) Delete(ctx context.Context, id int) error {
	result, err := r.DB.ExecContext(ctx, "DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.ErrEntityNotFound
	}
	return nil
}
