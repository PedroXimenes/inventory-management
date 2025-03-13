package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model

	WarehouseId string `json:"warehouse_id"`
	ProductId   string `json:"product_id"`
	Quantity    int64  `json:"quantity"`
}
