package handlers

import (
	"fmt"
	"inventory-management/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h handler) AddWarehouse(c *fiber.Ctx) error {

	warehouse := &models.Warehouses{}

	err := c.BodyParser(warehouse)
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.Status(http.StatusInternalServerError).SendString("Could not decode request body")
	}

	fmt.Println(warehouse)

	result := h.DB.Create(&warehouse)
	if result.Error != nil {
		fmt.Printf("Unable to create new record: %v\n", err)
	}
	fmt.Printf("result: %#v\n", result)

	return c.Status(http.StatusCreated).SendString("criou")
}
