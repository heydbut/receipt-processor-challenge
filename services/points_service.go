package services

import (
	"math"
	"receiptprocessor/models"
	"strconv"
	"strings"
	"time"
)

// CalculatePoints calculates the points for a receipt based on the specified rules
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// Convert total to float
	totalAmount, _ := strconv.ParseFloat(receipt.Total, 64)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if totalAmount == float64(int(totalAmount)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(totalAmount*100, 25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions with length multiple of 3
	for _, item := range receipt.Items {
		descriptionLength := len(strings.TrimSpace(item.ShortDescription))
		if descriptionLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Rule 7: 10 points if the time is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if (purchaseTime.Hour() == 14 && purchaseTime.Minute() >= 0) || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points
}

// Helper function to count alphanumeric characters
func countAlphanumeric(s string) int {
	count := 0
	for _, c := range s {
		if isAlphanumeric(c) {
			count++
		}
	}
	return count
}

// Helper function to check if a rune is alphanumeric
func isAlphanumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9')
}
