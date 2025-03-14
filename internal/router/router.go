package router

import (
	"inventory-management/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app *fiber.App, db *gorm.DB) {

	h := handlers.New(db)

	app.Get("/warehouses", h.ListWarehouse)
	app.Get("/warehouses/:id", h.GetWarehouse)
	app.Delete("/warehouses/:id", h.DeleteWarehouse)
	app.Post("/warehouses", h.AddWarehouse)
}
