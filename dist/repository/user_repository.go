package repository

import component "github.com/myyrakle/gopring_example/dist/compoent"

// @Repository
type UserRepository struct {
	postgresComponent *component.PostgresComponent
}

func (c *UserRepository) CreateUser(name string) error {
	return nil
}

func GopringNewComponent_UserRepository(postgresComponent *component.PostgresComponent, ) *UserRepository {
	return &UserRepository{
		postgresComponent: postgresComponent,
	}
}


