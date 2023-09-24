package service

import (
	"github.com/myyrakle/gopring_example/src/dto"
	"github.com/myyrakle/gopring_example/src/repository"
)

// @Service
type UserService struct {
	userRepository *repository.UserRepository
}

func (c *UserService) DoSignup(request dto.SignupRequest) string {
	c.userRepository.CreateUser(request)
	return "success"
}

func (c *UserService) DoLogin(request dto.LoginRequest) dto.LoginResponse {
	user := c.userRepository.FindUserByID(request.Id)

	if user == nil {
		return dto.LoginResponse{Success: false}
	}

	type User struct {
		Id       string `db:"id"`
		Name     string `db:"name"`
		Password string `db:"password"`
	}
	data := User{}
	user.Scan(&data)

	if data.Password != request.Password {
		return dto.LoginResponse{Success: false}
	} else {
		return dto.LoginResponse{Success: true, AccessToken: "...."}
	}
}
