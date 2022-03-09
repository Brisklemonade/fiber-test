package main

import (
	"github.com/Brisklemonade/fiber-test/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    fullPath := "./apps/web/build/index.html"
    
    app := fiber.New()

    app.Use(cors.New())
    
    api.Welcome(*app)
    // Retrieve static files and serve them on index route
    app.Static("/", "./apps/web/build")
    // Get request to index | On request, serve index.html from build folder
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile(fullPath)
    })

    app.Listen("localhost:8080")
}
