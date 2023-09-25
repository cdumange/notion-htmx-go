package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cdumange/notion-htmx-go/models"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db,
	}
}

const categoryProjection = "id, title"

// GetCategory search the category by ID
// if not found, will return models.ErrCategoryNotFound.
func (r *CategoryRepository) GetCategory(ctx context.Context, ID uuid.UUID) (models.Category, error) {
	c, err := r.search(ctx, fmt.Sprintf("SELECT %s FROM categories WHERE id = $1", categoryProjection), ID)
	if err != nil {
		return models.Category{}, fmt.Errorf("error while retrieving category(%s): %w", ID, err)
	}

	if len(c) == 0 {
		return models.Category{}, models.ErrCategoryNotFound
	}

	return c[0], err
}

// GetAll will retrieve all categories.
func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	return r.search(ctx, fmt.Sprintf("SELECT %s FROM categories", categoryProjection))
}

func (r *CategoryRepository) search(ctx context.Context, sql string, args ...any) ([]models.Category, error) {
	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error while searching tasks  (%s) | args: (%v): %w", sql, args, err)
	}

	defer func() { _ = rows.Close() }()

	var ret []models.Category

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, fmt.Errorf("error while reading rows: %w", err)
		}

		var t models.Category

		if err = rows.Scan(&t.ID, &t.Title); err != nil {
			return nil, fmt.Errorf("error while scanning row: %w", err)
		}

		ret = append(ret, t)
	}

	return ret, rows.Err()
}

// CreateCategory creates a category.
func (r *CategoryRepository) CreateCategory(ctx context.Context, category models.Category) (uuid.UUID, error) {
	ID := uuid.New()

	_, err := r.db.ExecContext(ctx, fmt.Sprintf("INSERT INTO categories(%s) VALUES($1, $2)",
		categoryProjection), ID, category.Title)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error while creating category: %w", err)
	}

	return ID, nil
}
