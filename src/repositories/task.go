package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
)

// NewTaskRepository returns a TaskRepository.
func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

// TaskRepository is a postgres repository for tasks.
type TaskRepository struct {
	db *sql.DB
}

const taskProjection = "id, category_id, title, creation_date"

// GetTaskByCategory retrieves the tasks for a category.
func (r *TaskRepository) GetTaskByCategory(ctx context.Context, categoryID uuid.UUID) ([]models.Task, error) {
	return r.search(ctx, fmt.Sprintf("SELECT %s FROM tasks WHERE category_id = $1", taskProjection), categoryID)
}

func (r *TaskRepository) search(ctx context.Context, sql string, args ...any) ([]models.Task, error) {
	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error while searching tasks  (%s) | args: (%v): %w", sql, args, err)
	}

	defer func() { _ = rows.Close() }()

	var ret []models.Task

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, fmt.Errorf("error while reading rows: %w", err)
		}

		var t models.Task

		if err = rows.Scan(&t.ID, &t.CategoryID, &t.Title,
			&t.CreationDate); err != nil {
			return nil, fmt.Errorf("error while scanning row: %w", err)
		}

		ret = append(ret, t)
	}

	return ret, rows.Err()
}

// CreateTask creates a task.
func (r *TaskRepository) CreateTask(ctx context.Context, task models.Task) (uuid.UUID, error) {
	ID := uuid.New()

	_, err := r.db.ExecContext(ctx,
		fmt.Sprintf("INSERT INTO tasks(%s) VALUES($1, $2, $3, now())", taskProjection),
		ID, task.CategoryID, task.Title)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error creating task: %w", err)
	}

	return ID, nil
}

// UpdateTask updates a task.
func (r *TaskRepository) UpdateTask(ctx context.Context, task models.Task) error {
	_, err := r.db.ExecContext(ctx, "UPDATE tasks set title= $1 WHERE id=$2", task.Title, task.ID)
	return err
}

// DeleteTask deletes a task by its ID.
func (r *TaskRepository) DeleteTask(ctx context.Context, ID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tasks WHERE id=$1", ID)
	return err
}
