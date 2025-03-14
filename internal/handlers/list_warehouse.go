package handlers

import (
	"fmt"
	"inventory-management/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h handler) ListWarehouse(c *fiber.Ctx) error {

	warehouses := &[]models.Warehouses{}

	if result := h.DB.Find(&warehouses); result.Error != nil {
		fmt.Println(result.Error)
	}

	return c.Status(http.StatusOK).JSON(warehouses)
}
