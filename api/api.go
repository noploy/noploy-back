package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noploy/noploy-back/api/v1/auth"
)

func NewAPI(app *fiber.App) {
	auth.NewAuth(app)
}
