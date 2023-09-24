package repository

import (
	component "github.com/myyrakle/gopring_example/dist/compoent"
	"github.com/myyrakle/gopring_example/dist/dto"
	"github.com/myyrakle/gopring_example/dist/model"
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

func (c *UserRepository) FindUserByID(id string) (model.User, error) {
	db := c.postgresComponent.DB()
	rows, err := db.Query("SELECT id, name, password FROM users WHERE id = $1", id)

	if err != nil {
		return model.User{}, err
	}

	_id := ""
	name := ""
	password := ""
	rows.Next()
	rows.Scan(&_id, &name, &password)

	return model.User{
		Id:       _id,
		Name:     name,
		Password: password,
	}, nil
}

func GopringNewComponent_UserRepository(postgresComponent *component.PostgresComponent, ) *UserRepository {
	return &UserRepository{
		postgresComponent: postgresComponent,
	}
}


