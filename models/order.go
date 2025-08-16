package models

import "time"

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"user_id"`
	User       User        `json:"user"`
	Status     string      `json:"status" gorm:"default:pending"` // pending|paid|shipped|completed|cancelled
	TotalPrice int64       `json:"total_price"`
	Items      []OrderItem `json:"items" gorm:"constraint:OnDelete:CASCADE;"`
	ProofURL   string      `json:"proof_url"` // bukti pembayaran (opsional)
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
