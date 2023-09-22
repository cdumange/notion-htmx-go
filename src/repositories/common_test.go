package repositories

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func cleanTable(t *testing.T, db *sql.DB) {
	t.Helper()

	_, err := db.Exec("DELETE FROM tasks;")
	require.NoError(t, err)
}

func initDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost/public?sslmode=disable")
	require.NoError(t, err)
	cleanTable(t, db)
	return db
}
