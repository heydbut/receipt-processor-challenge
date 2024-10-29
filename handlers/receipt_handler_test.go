package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"receiptprocessor/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

import (
	"bytes"
	"encoding/json"
	"receiptprocessor/handlers"
)

func TestProcessReceipt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handlers.SetupRoutes(router)

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

	jsonData, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "id")
}

func TestGetPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handlers.SetupRoutes(router)

	// First, process a receipt to get an ID
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

	jsonData, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	receiptID := response["id"]

	// Now, retrieve the points using the ID
	req, _ = http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var pointsResponse map[string]int
	err := json.Unmarshal(w.Body.Bytes(), &pointsResponse)
	assert.NoError(t, err)
	assert.Equal(t, 31, pointsResponse["points"])
}

func TestGetPoints_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handlers.SetupRoutes(router)

	req, _ := http.NewRequest("GET", "/receipts/invalid-id/points", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
