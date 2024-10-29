package services_test

import (
	"receiptprocessor/models"
	"receiptprocessor/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints_SimpleReceipt(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        "1.25",
		Items: []models.Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            "1.25",
			},
		},
	}

	points := services.CalculatePoints(receipt)
	assert.Equal(t, 31, points)
}

func TestCalculatePoints_MorningReceipt(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items: []models.Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            "1.25",
			},
			{
				ShortDescription: "Dasani",
				Price:            "1.40",
			},
		},
	}

	points := services.CalculatePoints(receipt)
	assert.Equal(t, 15, points)
}

func TestCalculatePoints_ComplexReceipt(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []models.Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}

	points := services.CalculatePoints(receipt)
	assert.Equal(t, 109, points)
}
