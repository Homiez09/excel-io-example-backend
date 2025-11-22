package models

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"size:50;uniqueIndex" json:"code"`
	Name      string    `gorm:"size:255" json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}
