package entity

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:varchar(255);not null"`
	Price       float64 `gorm:"type:decimal(15,2);not null"`
	Stock       int     `gorm:"not null"`
	Category    string  `gorm:"type:varchar(100);not null"`
}

type ProductsResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
}

type ProductsRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       *int    `json:"stock"`
	Category    string  `json:"category"`
}
