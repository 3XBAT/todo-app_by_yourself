package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	// Драйвер для выполнения миграций Postgres
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// Драйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://./schema",
		fmt.Sprintf("postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable"),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}

	fmt.Println("Applied migrations")
}
