package models

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a task.
type Task struct {
	ID         uuid.UUID `json:"id"`
	CategoryID uuid.UUID `form:"category_id" json:"category_id" validate:"required"`
	Title      string    `form:"title" json:"title" validate:"required"`

	CreationDate time.Time `json:"creation_date"`
}

// Category represent a task category.
type Category struct {
	ID    uuid.UUID
	Title string
	List  []Task
}
