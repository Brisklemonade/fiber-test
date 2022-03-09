package api

import (
	"github.com/gofiber/fiber/v2"
)

/**
* JSON object that denotes a status, and sends a message
 */
func Welcome(app fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error {
        
        return c.JSON(&fiber.Map{
            "success": true,
            "message": "welcome",
        })
    })
}