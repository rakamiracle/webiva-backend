package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"` // simpan dalam satuan terkecil (mis. rupiah)
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
