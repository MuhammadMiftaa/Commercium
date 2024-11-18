package entity

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique;not null"`
	Fullname string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Income   int    `gorm:"type:decimal(10,2)"`
	Outcome  int    `gorm:"type:decimal(10,2)"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     string `gorm:"type:varchar(100);not null"`
}

type UsersResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Income   int    `json:"income"`
	Outcome  int    `json:"outcome"`
	Role     string `json:"role"`
}

type UsersRequest struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Income   int    `json:"income"`
	Outcome  int    `json:"outcome"`
	Password string `json:"password"`
}
