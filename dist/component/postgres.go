package compoent

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// Component
type PostgresComponent struct {
}

func (c *PostgresComponent) Init() {
	var err error

	connStr := fmt.Sprintf("user=pqgotest dbname=pqgotest sslmode=verify-full")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}


