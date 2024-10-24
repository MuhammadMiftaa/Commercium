package entity

import "time"

type Products struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Stock       int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null"`
}

type ProductsResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type ProductsRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       *int    `json:"stock"`
}
