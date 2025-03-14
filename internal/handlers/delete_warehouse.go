package handlers

import (
	"fmt"
	"inventory-management/internal/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteWarehouse(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' is required")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' must be an integer")
	}

	var warehouse models.Warehouses

	if result := h.DB.Delete(&warehouse, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	return c.Status(http.StatusOK).JSON(warehouse)
}
