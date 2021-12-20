package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noploy/noploy-back/common"
)

type API struct {
}

func (a *API) Login(ctx *fiber.Ctx) error {
	return ctx.JSON(common.JSON{"success": true})
}

func NewAuth(app *fiber.App) {
	var api API
	authGroup := app.Group("/auth")
	{
		authGroup.Post("/login", api.Login)
	}
}
