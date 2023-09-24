package main

import (
	"os"

	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	fx "go.uber.org/fx" 
	"github.com/myyrakle/gopring/pkg/properties"
	gp000003 "github.com/myyrakle/gopring_example/dist/controller"
	gp000004 "github.com/myyrakle/gopring_example/dist/dto"
	gp000005 "github.com/myyrakle/gopring_example/dist/dto"
	gp000008 "github.com/myyrakle/gopring_example/dist/repository"
	gp000009 "github.com/myyrakle/gopring_example/dist/service"
	gp000002 "github.com/myyrakle/gopring_example/dist/compoent"

)

func RunGopring(appProperties *properties.Properties, gc000001 *gp000003.HomeController, gc000002 *gp000003.UserController, ) {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {

		response := gc000001.Index(c, )

		return c.JSON(200, response)
	})
	app.GET("/health", func(c echo.Context) error {

		response := gc000001.HelathCheck(c, )

		return c.JSON(200, response)
	})
	app.POST("/user/signup", func(c echo.Context) error {
		body := gp000004.SignupRequest{}
		if err := c.Bind(&body); err != nil {
			return err
		}
		response := gc000002.Signup(c, body)

		return c.JSON(200, response)
	})
	app.POST("/user/login", func(c echo.Context) error {
		body := gp000005.LoginRequest{}
		if err := c.Bind(&body); err != nil {
			return err
		}
		response := gc000002.Login(c, body)

		return c.JSON(200, response)
	})


	port := ":8080"

	if appProperties != nil {
		property := properties.FindByKey(*appProperties, "server.port")
		if property != nil {
			port = ":" + property.Value
		}
	}

	app.Logger.Fatal(app.Start(port))

	app.Logger.Fatal(app.Start(":8080"))
}

func LoadApplicationProperties() *properties.Properties {
	if _, err := os.Stat("application.properties"); os.IsNotExist(err) {
		return &properties.Properties{}
	} else {
		if fileData, err := os.ReadFile("application.properties"); err != nil {
			panic(err)
		} else {
			parsedProperties := properties.ParseProperties(string(fileData))
			return &parsedProperties
		}
	}
}

func main() {
	providers := fx.Provide(
		LoadApplicationProperties,
		gp000002.GopringNewComponent_PostgresComponent,
		gp000003.GopringNewController_HomeController,
		gp000003.GopringNewController_UserController,
		gp000008.GopringNewComponent_UserRepository,
		gp000009.GopringNewComponent_UserService,
		
	)

	fx.New(providers, fx.Invoke(RunGopring)).Run()
}
