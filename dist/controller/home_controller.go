package controller

import (
	"github.com/labstack/echo"
)

// @Controller(/)
type HomeController struct {
}

// @GetMapping("/")
func (this HomeController) Index(c echo.Context) string {
	return "hello world"
}

type HealthCheckResponse struct {
	Ok bool `json:"ok"`
}

// @GetMapping("/health")
func (this *HomeController) HelathCheck(c echo.Context) HealthCheckResponse {
	return HealthCheckResponse{
		Ok: true,
	}
}


func GopringNewController_HomeController() *HomeController {
	return &HomeController{
	}
}

