package main

import (
	"fiber-base-go/database"
	"fiber-base-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	app.Listen(":3000")
}
