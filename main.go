package main

import (
	"github.com/Brisklemonade/fiber-test/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()
    
    /** Allow cross origin & things */
    app.Use(cors.New())
    
    /** Initialize the api */
    api.Welcome(*app)
    
    /**
    * Retrieve static files and serve them on index route
    * Get request to index | On request, serve index.html from build folder
    */
    fullPath := "./apps/web/build/index.html"
    app.Static("/", "./apps/web/build")
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile(fullPath)
    })

    app.Listen("localhost:8080")
}
