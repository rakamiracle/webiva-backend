package models

import "time"

type StoreSetting struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	StoreName       string    `json:"store_name"`
	LogoURL         string    `json:"logo_url"`
	PaymentMethods  string    `json:"payment_methods"` // CSV sederhana: "transfer,qris,cod"
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
