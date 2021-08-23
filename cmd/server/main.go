package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tocoteron/toco-auth/handler"
	custom_middleware "github.com/tocoteron/toco-auth/middleware"
	"github.com/tocoteron/toco-auth/model"
)

func main() {
	identifier := flag.String("identifier", "", "Identifier")
	secret := flag.String("secret", "", "Secret key")
	flag.Parse()

	if *identifier == "" {
		fmt.Println("--identifier option is required")
		return
	}

	if *secret == "" {
		fmt.Println("--secret option is required")
		return
	}

	setting := &model.ServerSetting{
		Identifier: *identifier,
		Secret:     *secret,
	}

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)

	e.Logger.Infof("[Server setting] Identifier: %s, Secret key: %s", *identifier, *secret)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	authGroup := e.Group("/auth", custom_middleware.ServerSettingProvider(setting))
	authGroup.POST("/signup", handler.SignUp)
	authGroup.POST("/signin", handler.SignIn)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
