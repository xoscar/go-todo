package connectors

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB               *sql.DB
	connectionString string
	migrationsFolder string
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func NewPostgres(connectionString string, migrationsFolder string) (Postgres, error) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return Postgres{}, fmt.Errorf("open: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return Postgres{}, fmt.Errorf("with instance: %w", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(migrationsFolder, "postgres", driver)

	if err != nil {
		return Postgres{}, fmt.Errorf("migrate new: %w", err)
	}

	migrations.Up()

	return Postgres{connectionString: connectionString, migrationsFolder: migrationsFolder, DB: db}, nil
}
