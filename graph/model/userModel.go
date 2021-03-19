package model

import "time"

type User struct {
	ID         int         `json:"id"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Role       string      `json:"role"`
	Whatsapp   string      `json:"whatsapp"`
	Password   string      `json:"password"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Avatar     *int        `json:"avatar"`
	Image      Image       `gorm:"-"`
	Friends    []*User     `gorm:"-"`
	LookingFor []string    `gorm:"-"`
	Comodity   []*Comodity `gorm:"-"`
}
