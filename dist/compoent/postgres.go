package component

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/myyrakle/gopring/pkg/properties"
)

var db *sql.DB

// @Component
type PostgresComponent struct {
	appProperties *properties.Properties
}

func (c *PostgresComponent) Init() {
	var err error

	connStr := fmt.Sprintf("user=pqgotest dbname=pqgotest sslmode=verify-full")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func GopringNewComponent_PostgresComponent(appProperties *properties.Properties, ) *PostgresComponent {
	return &PostgresComponent{
		appProperties: appProperties,
	}
}


