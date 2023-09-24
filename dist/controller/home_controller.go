package controller

import (
	"github.com/labstack/echo"
	"github.com/myyrakle/gopring_example/dist/service"
)

// @Controller(/)
type HomeController struct {
	service *service.HomeService
}

// @GetMapping("/")
func (this HomeController) Index(c echo.Context) string {
	return this.service.GetHello()
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


func GopringNewController_HomeController(service *service.HomeService, ) *HomeController {
	return &HomeController{
		service: service,
	}
}

