package models

import (
	"github.com/cdumange/notion-htmx-go/common"
	"github.com/google/uuid"
)

// ErrCategoryNotFound is an error triggered if the researched category was not found.
const ErrCategoryNotFound = common.Error("category was not found")

// Category represent a task category.
type Category struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title" form:"title"`
	List  []Task    `json:"list"`
}
