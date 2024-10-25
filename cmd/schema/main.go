package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	root "github.com/10ego/gthp"
	"github.com/pressly/goose/v3"
)

func main() {
	db_url := os.Getenv("DATABASE_URL")
	// setup database

	// Register the PostgreSQL driver
	if _, err := sql.Open("postgres", ""); err != nil {
		fmt.Println("Failed to find driver")
		panic(err)
	}

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}
	defer db.Close()

	goose.SetBaseFS(root.GetMigrationFS())

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	fmt.Println("Migrations ran successfully!")
	// run app
}
