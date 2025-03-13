package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model

	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}
