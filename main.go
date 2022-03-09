package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
    fullPath := "./apps/web/client/index.html"
    
    app := fiber.New()

    app.Get("/api", func(c *fiber.Ctx) error {
        
        return c.JSON(&fiber.Map{
            "success": true,
            "message": "welcome",
        })
    })
    
    app.Static("/", "./apps/web/client")

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile(fullPath)
    })

    app.Listen("localhost:3000")
}
