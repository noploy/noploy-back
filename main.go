package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/noploy/noploy-back/api"
	"github.com/sirupsen/logrus"
)

func main() {
	var env = LoadEnvironment()
	app := fiber.New()
	api.NewAPI(app)
	err := app.Listen(fmt.Sprintf(":%s", env.Port))
	if err != nil {
		logrus.WithField("error", err.Error()).Error("Error in listen server.")
	}
}
