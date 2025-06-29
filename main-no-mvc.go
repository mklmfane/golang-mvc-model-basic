package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // use IPv4 explicitly
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	// Build connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// CRUD examples with timing
	createAccount(db, 1, "Alice Smith")
	getAccount(db, 1)
	updateAccount(db, 1, "Alice Johnson")
	deleteAccount(db, 1)
}

// Create
func createAccount(db *sql.DB, id int, fullName string) {
	start := time.Now()

	sqlStatement := `INSERT INTO "accounts" ("id", "full_name") VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, id, fullName)
	if err != nil {
		log.Println("Error inserting:", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("New account created. Took %s\n", elapsed)
}

// Read
func getAccount(db *sql.DB, id int) {
	start := time.Now()

	var fullName string
	sqlStatement := `SELECT "full_name" FROM "accounts" WHERE "id" = $1`
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&fullName); err {
	case sql.ErrNoRows:
		fmt.Println("No account found.")
	case nil:
		fmt.Printf("Account %d: %s\n", id, fullName)
	default:
		log.Println("Error reading:", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Query took %s\n", elapsed)
}

// Update
func updateAccount(db *sql.DB, id int, newName string) {
	start := time.Now()

	sqlStatement := `UPDATE "accounts" SET "full_name" = $2 WHERE "id" = $1`
	res, err := db.Exec(sqlStatement, id, newName)
	if err != nil {
		log.Println("Error updating:", err)
		return
	}
	count, _ := res.RowsAffected()
	fmt.Printf("Updated %d rows.\n", count)

	elapsed := time.Since(start)
	fmt.Printf("Update took %s\n", elapsed)
}

// Delete
func deleteAccount(db *sql.DB, id int) {
	start := time.Now()

	sqlStatement := `DELETE FROM "accounts" WHERE "id" = $1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Println("Error deleting:", err)
		return
	}
	count, _ := res.RowsAffected()
	fmt.Printf("Deleted %d rows.\n", count)

	elapsed := time.Since(start)
	fmt.Printf("Delete took %s\n", elapsed)
}
