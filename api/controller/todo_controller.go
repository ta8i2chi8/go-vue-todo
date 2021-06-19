package controller

import (
	"net/http"
	"strconv"
	"time"

	"go-vue-todo/usecase"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	UseCase *usecase.TodoUseCase
}

type ListParams struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Context   string `json:"context"`
	LimitDate string `json:"limit_date"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (c *TodoController) List(ctx *gin.Context) {
	queryName, _ := ctx.GetQuery("title")
	todos, err := c.UseCase.List(ctx, queryName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	const format = "2006-01-02 15:04:05"

	outs := make([]*ListParams, 0)

	for i := 0; i < len(todos); i++ {
		out := &ListParams{
			ID:        todos[i].ID,
			Title:     todos[i].Title,
			Context:   todos[i].Context,
			LimitDate: todos[i].LimitDate.Format(format),
			CreatedAt: todos[i].CreatedAt.Format(format),
			UpdatedAt: todos[i].UpdatedAt.Format(format),
		}

		outs = append(outs, out)
	}

	ctx.JSON(http.StatusOK, outs)
}

func (c *TodoController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todo, err := c.UseCase.GetById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewError(err))
	}

	const format = "2006-01-02 15:04:05"

	out := &ListParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Context:   todo.Context,
		LimitDate: todo.LimitDate.Format(format),
		CreatedAt: todo.CreatedAt.Format(format),
		UpdatedAt: todo.UpdatedAt.Format(format),
	}

	ctx.JSON(http.StatusOK, out)
}

func (c *TodoController) Create(ctx *gin.Context) {
	params := &struct {
		Title     string `json:"title"`
		Context   string `json:"context"`
		LimitDate string `json:"limit_date"`
	}{}

	if err := ctx.BindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err))
		return
	}

	lct, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", params.LimitDate, lct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err))
		return
	}

	in := &usecase.CreateParams{
		Title:     params.Title,
		Context:   params.Context,
		LimitDate: t,
	}

	if err := c.UseCase.Create(ctx, in); err != nil {
		ctx.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create success!",
	})
}

func (c *TodoController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	params := &struct {
		Title     string `json:"title"`
		Context   string `json:"context"`
		LimitDate string `json:"limit_date"`
	}{}
	if err := ctx.BindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err))
		return
	}

	lct, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", params.LimitDate, lct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err))
		return
	}

	in := &usecase.UpdateParams{
		ID:        id,
		Title:     params.Title,
		Context:   params.Context,
		LimitDate: t,
	}
	if err := c.UseCase.Update(ctx, in); err != nil {
		ctx.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update success!",
	})
}

func (c *TodoController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.UseCase.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete success!",
	})
}