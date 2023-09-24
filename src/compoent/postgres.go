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

func (this *PostgresComponent) Init() {
	if db == nil {
		var err error

		host := properties.FindByKey(*this.appProperties, "db.host").Value
		username := properties.FindByKey(*this.appProperties, "spring.datasource.username").Value
		password := properties.FindByKey(*this.appProperties, "spring.datasource.password").Value
		dbname := properties.FindByKey(*this.appProperties, "spring.datasource.dbname").Value

		connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, host, dbname)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (this *PostgresComponent) DB() *sql.DB {
	if db == nil {
		this.Init()
	}

	return db
}
