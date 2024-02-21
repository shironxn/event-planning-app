package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"uniqueIndex"`
	Email    string `json:"email" form:"email" gorm:"uniqueIndex"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" form:"name" validate:"required,min=4,max=30"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserDetails struct {
	Token   string `json:"token,omitempty"`
	Expired string `json:"expired_at,omitempty"`
}
