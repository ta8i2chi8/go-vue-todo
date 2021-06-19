package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"         db:"id"`
	Title     string    `json:"title"      db:"title"`
	Context   string    `json:"context"    db:"context"`
	LimitDate time.Time `json:"limit_date" db:"limitDate"`
	CreatedAt time.Time `json:"created_at" db:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" db:"updatedAt"`
}