package entity

import "time"

type Orders struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	ProductID  int       `gorm:"not null"`
	Quantity   int       `gorm:"not null"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null"`
	Status     string    `gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null"`
}
