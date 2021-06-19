package main

import (
    "time"

    "go-vue-todo/controller"
    "go-vue-todo/infra"
    "go-vue-todo/usecase"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // dsn := "root:@tcp(127.0.0.1:3306)/todo-db?parseTime=true"
    dsn := "root:@tcp(todo-mysql)/todo-db?parseTime=true&loc=Asia%2FTokyo"
    db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        panic(err)
    }
    db.SetConnMaxLifetime(3 * time.Minute)
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(10)

    todoRepo := &infra.MySQLTodoRepository{db}

    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowHeaders: []string{
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "X-CSRF-Token",
            "Authorization",
        },
        AllowOrigins: []string{"*"},
        MaxAge: 24 * time.Hour,
    }))

    todo := &controller.TodoController{
        UseCase: usecase.NewTodoUseCase(todoRepo),
    }

    r.GET("/todos", todo.List)
    r.GET("/todos/:id", todo.GetById)
    r.POST("/todos", todo.Create)
    r.PUT("todos/:id", todo.Update)
    r.DELETE("todos/:id", todo.Delete)

    r.Run(":8888")
}