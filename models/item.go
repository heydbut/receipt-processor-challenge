package models

// Item represents an item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required,shortDescriptionPattern"`
	Price            string `json:"price" binding:"required,pricePattern"`
}
