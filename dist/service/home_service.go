package service

import "github.com/myyrakle/gopring_example/dist/repository"

// @Service
type HomeService struct {
	userRepository *repository.UserRepository
}

func (c *HomeService) GetHello() string {
	c.userRepository.CreateUser("test")
	return "Hello World!"
}

func GopringNewComponent_HomeService(userRepository *repository.UserRepository, ) *HomeService {
	return &HomeService{
		userRepository: userRepository,
	}
}


