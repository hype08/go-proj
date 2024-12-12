package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp(db *DB, migrationDir string) error {
	driver, err := postgres.WithInstance(db.Pool.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = migration.Up()

	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func MigrateDown(db *DB, migrationDir string) error {
	driver, err := postgres.WithInstance(db.Pool.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = migration.Down()

	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}
