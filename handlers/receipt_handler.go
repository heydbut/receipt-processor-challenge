package handlers

import (
	"net/http"
	"receiptprocessor/models"
	"receiptprocessor/services"
	"receiptprocessor/utils"
	"sync"

	"github.com/gin-gonic/gin"
)

// Use sync.Map for thread-safe operations
var receiptsDB = sync.Map{}

// SetupRoutes initializes the API routes
func SetupRoutes(router *gin.Engine) {
	utils.SetupValidation()

	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)
}

// Handler to process receipts
func processReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receiptID := receipt.GenerateID()

	points := services.CalculatePoints(receipt)

	receiptsDB.Store(receiptID, points)

	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

// Handler to get points for a receipt
func getPoints(c *gin.Context) {
	id := c.Param("id")

	value, exists := receiptsDB.Load(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}

	points, ok := value.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid data type for points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
