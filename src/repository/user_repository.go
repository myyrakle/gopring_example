package repository

import (
	"database/sql"

	component "github.com/myyrakle/gopring_example/src/compoent"
	"github.com/myyrakle/gopring_example/src/dto"
)

// @Repository
type UserRepository struct {
	postgresComponent *component.PostgresComponent
}

func (c *UserRepository) CreateUser(request dto.SignupRequest) error {
	db := c.postgresComponent.DB()
	_, err := db.Exec("INSERT INTO users (id, name, password) VALUES ($1, $2, $3)", request.Id, request.Name, request.Password)

	return err
}

func (c *UserRepository) FindUserByID(id string) *sql.Row {
	db := c.postgresComponent.DB()
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	return row
}
