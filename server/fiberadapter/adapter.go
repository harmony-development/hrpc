package fiberadapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/hrpc/server"
)

func RegisterFiber(serviceHandler server.HRPCServiceHandler, app *fiber.App) {
	for path, handler := range serviceHandler.Routes() {
		app.All(path, func(c *fiber.Ctx) error {
			resp, err := handler(c.Context(), c.Request())
			if err != nil {
				return err
			}
			return c.Send(resp)
		})
	}
}
