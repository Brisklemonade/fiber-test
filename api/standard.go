package api

import (
	"github.com/gofiber/fiber/v2"
)

/**
* This does legit nothing still figuring things out
 */
func Welcome(c *fiber.Ctx) interface {} {
	welcome := "Welcome to the api"

	return c.JSON(&fiber.Map{
			"success": true,
			"message": welcome,
	})
}