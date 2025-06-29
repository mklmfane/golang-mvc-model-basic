package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/internal/controller"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // ensure IPv4
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Cannot open db:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// Automatically create the table if it doesn't exist
	if err := createTableIfNotExists(db); err != nil {
		log.Fatal("Failed creating table:", err)
	}
	fmt.Println("Ensured accounts_bank table exists.")

	// Now run your MVC operations
	accController := controller.NewAccountController(db)

	accController.CreateAccount("Alice Smith")
	accController.GetAccount(1)
	accController.UpdateAccount(1, "Alice Johnson")
	accController.DeleteAccount(1)
}

// Automatically create the accounts_bank table if it doesn't exist
func createTableIfNotExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS accounts_bank (
		id SERIAL PRIMARY KEY,
		full_name VARCHAR NOT NULL,
		created_at TIMESTAMPTZ DEFAULT now()
	);
	`
	_, err := db.Exec(query)
	return err
}
