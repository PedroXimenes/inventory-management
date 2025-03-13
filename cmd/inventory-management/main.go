package main

import (
	"inventory-management/internal/db"
	"inventory-management/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	db := db.Init()
	router.Router(app, db)

	err := app.Listen(":9000")
	if err != nil {
		log.Fatalf("Could not start server at port: %s\nError: %s", ":9000", err.Error())
	}
}
