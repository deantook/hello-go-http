package model

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customer_name" binding:"required"`
	ProductName  string    `json:"product_name" binding:"required"`
	Amount       float64   `json:"amount" binding:"required,gt=0"`
	Status       string    `json:"status" gorm:"default:pending"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
	CustomerName string  `json:"customer_name" binding:"required"`
	ProductName  string  `json:"product_name" binding:"required"`
	Amount       float64 `json:"amount" binding:"required,gt=0"`
	Status       string  `json:"status"`
}

type UpdateOrderRequest struct {
	CustomerName *string  `json:"customer_name"`
	ProductName  *string  `json:"product_name"`
	Amount       *float64 `json:"amount" binding:"omitempty,gt=0"`
	Status       *string  `json:"status"`
}
