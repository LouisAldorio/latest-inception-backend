package model

import "time"

type Comodity struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Image       []*string  `gorm:"-"`
	UnitPrice   float64    `json:"unit_price"`
	UnitType    string     `json:"unit_type"`
	MinPurchase int        `json:"min_purchase"`
	Description *string    `json:"description"`
	UserID      int        `json:"user_id"`
	CategoryID  int        `json:"category_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	User        *User      `gorm:"-"`
}
