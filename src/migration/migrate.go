package migration

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(databaseURL string) error {
	m, err := migrate.New(
		"file://./db/migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("migrate.New error: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate.Up error: %w", err)
	}

	log.Printf("migrate")

	return nil
}
