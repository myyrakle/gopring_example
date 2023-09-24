package controller

import (
	"github.com/labstack/echo"
	"github.com/myyrakle/gopring_example/src/dto"
	"github.com/myyrakle/gopring_example/src/service"
)

// @Controller("/user")
type UserController struct {
	service *service.UserService
}

// @PostMapping("/signup")
func (this *UserController) Signup(
	c echo.Context,
	// @RequestBody
	body dto.SignupRequest,
) string {
	return this.service.DoSignup(body)
}

// @PostMapping("/login")
func (this *UserController) Login(
	c echo.Context,
	// @RequestBody
	body dto.LoginRequest,
) dto.LoginResponse {
	return this.service.DoLogin(body)
}
