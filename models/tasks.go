package models

import "github.com/google/uuid"

type Task struct {
	ID uuid.UUID 
	CategoryID uuid.UUID `form:"category_id"`
	Title string `form:"title"`
}

type TaskList struct {
	Tasks []Task
}

type Categories struct {
	ID uuid.UUID
	Title string
	List TaskList
}