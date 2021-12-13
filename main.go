package main

import (
	"github.com/KenMwaura1/go-todo-api/database"
	"github.com/KenMwaura1/go-todo-api/todo"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
import "fmt"

func main() {
	fmt.Println("App running!")
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":5002"))
}
