package entity

import "time"

type Users struct {
	ID        int
	Username  string    `gorm:"type:varchar(100);unique;not null"`
	Fullname  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Role      string    `gorm:"type:varchar(100);not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type UsersResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UsersRequest struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
