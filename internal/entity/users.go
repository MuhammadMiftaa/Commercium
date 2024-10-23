package entity

import "time"

type Users struct {
	ID        int
	Username  string    `gorm:"type:varchar(100);unique;not null"`
	Fullname  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Role      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}
