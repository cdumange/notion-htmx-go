package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

type postgresSuite struct {
	suite.Suite
	db *sql.DB
}

func Test_Suite(t *testing.T) {
	suite.Run(t, new(postgresSuite))
}

func (s *postgresSuite) SetupSuite() {
	t := s.T()
	s.db = initDockerDB(t)
	deployMigrations(t, s.db)
}

func initDockerDB(t *testing.T) *sql.DB {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	require.NoError(t, pool.Client.Ping())

	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       "notiontests",
		Repository: "postgres",
		Tag:        "alpine",
		Env: []string{
			"POSTGRES_USERNAME=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=public",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{
					HostIP:   "",
					HostPort: "8000",
				},
			},
		},
		Privileged: false,
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := pool.Purge(container); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	})

	var db *sql.DB

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres",
			fmt.Sprintf(
				"postgresql://postgres:postgres@%s/public?sslmode=disable",
				container.GetHostPort("5432/tcp")),
		)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	return db
}

func deployMigrations(t *testing.T, db *sql.DB) {
	t.Helper()

	folder, err := os.Open("../migration/sql")
	require.NoError(t, err)

	files, err := folder.ReadDir(-1)
	require.NoError(t, err)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		req, err := os.ReadFile("../migration/sql/" + file.Name())
		require.NoError(t, err)

		_, err = db.Exec(string(req))
		require.NoError(t, err)
	}
}
