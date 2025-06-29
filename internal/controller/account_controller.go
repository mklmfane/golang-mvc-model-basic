package controller

import (
	"context"
	"fmt"
	"time"

	"database/sql"

	"myapp/internal/model"
)

type AccountController struct {
	queries *model.Queries
	db      *sql.DB
}

func NewAccountController(db *sql.DB) *AccountController {
	return &AccountController{
		queries: model.New(db),
		db:      db,
	}
}

func (c *AccountController) CreateAccount(fullName string) {
	start := time.Now()
	account, err := c.queries.CreateAccount(context.Background(), fullName)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	fmt.Printf("Created: %+v (took %s)\n", account, time.Since(start))
}

func (c *AccountController) GetAccount(id int64) {
	start := time.Now()
	account, err := c.queries.GetAccount(context.Background(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No account found.")
			return
		}
		fmt.Println("Error fetching account:", err)
		return
	}
	fmt.Printf("Fetched: %+v (took %s)\n", account, time.Since(start))
}

func (c *AccountController) UpdateAccount(id int64, newName string) {
	start := time.Now()
	account, err := c.queries.UpdateAccount(context.Background(), model.UpdateAccountParams{
		ID:       int32(id),
		FullName: newName,
	})
	if err != nil {
		fmt.Println("Error updating account:", err)
		return
	}
	fmt.Printf("Updated: %+v (took %s)\n", account, time.Since(start))
}

func (c *AccountController) DeleteAccount(id int64) {
	start := time.Now()
	err := c.queries.DeleteAccount(context.Background(), int32(id))
	if err != nil {
		fmt.Println("Error deleting account:", err)
		return
	}
	fmt.Printf("Deleted account id=%d (took %s)\n", id, time.Since(start))
}
