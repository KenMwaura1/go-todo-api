package main 

import (
	"log"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)
import "fmt"

func main() {
	fmt.Println("App running!")
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":5002"))
}
