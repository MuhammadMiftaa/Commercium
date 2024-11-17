package entity

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID     int     `gorm:"not null"`
	ProductID  int     `gorm:"not null"`
	Quantity   int     `gorm:"not null"`
	TotalPrice float64 `gorm:"type:decimal(15,2)"`
	Status     string  `gorm:"type:varchar(100);not null"`

	User    Users    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product Products `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type OrderDetail struct {
	ID           uint    `json:"id"`
	CustomerName string  `json:"customer_name"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	ProductPrice float64 `json:"product_price"`
	TotalPrice   float64 `json:"total_price"`
	Status       string  `json:"status"`
}

type OrdersResponse struct {
	ID         uint    `json:"id"`
	UserID     int     `json:"user_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalprice"`
	Status     string  `json:"status"`
}

type OrdersRequest struct {
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}
