package models

import (
	"github.com/google/uuid"
)

// Receipt represents the structure of a receipt
type Receipt struct {
	Retailer     string `json:"retailer" binding:"required,retailerPattern"`
	PurchaseDate string `json:"purchaseDate" binding:"required,purchaseDatePattern"`
	PurchaseTime string `json:"purchaseTime" binding:"required,purchaseTimePattern"`
	Items        []Item `json:"items" binding:"required,min=1,dive"`
	Total        string `json:"total" binding:"required,totalPattern"`
}

// GenerateID generates a new UUID for the receipt
func (r *Receipt) GenerateID() string {
	return uuid.New().String()
}
