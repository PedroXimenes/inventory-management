package router

import (
	"inventory-management/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {

	h := handlers.New(db)

	app.Post("/warehouse", h.AddWarehouse)
}
