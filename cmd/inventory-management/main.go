package main

import (
	"inventory-management/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	router.Router(app)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Could not start server at port: %s\nError: %s", "8080", err.Error())
	}
}
