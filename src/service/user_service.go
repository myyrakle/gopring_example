package service

import (
	"fmt"

	"github.com/myyrakle/gopring_example/src/dto"
	"github.com/myyrakle/gopring_example/src/repository"
)

// @Service
type UserService struct {
	userRepository *repository.UserRepository
}

func (c *UserService) DoSignup(request dto.SignupRequest) string {
	if err := c.userRepository.CreateUser(request); err != nil {
		fmt.Println(err)
		return "fail"
	}
	return "success"
}

func (c *UserService) DoLogin(request dto.LoginRequest) dto.LoginResponse {
	user, err := c.userRepository.FindUserByID(request.Id)

	if err != nil {
		fmt.Println(err)
		return dto.LoginResponse{Success: false}
	}

	if user.Password != request.Password {
		return dto.LoginResponse{Success: false}
	} else {
		return dto.LoginResponse{Success: true, AccessToken: "...."}
	}
}
