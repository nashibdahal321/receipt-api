package calculator

import (
	"receipt-api/model"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	// Create a sample receipt for testing
	receipt := &model.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []model.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}

	// Call the function you want to test
	points := CalculatePoints(receipt)

	// Add assertions to validate the result
	expectedPoints := 6 + 10 + 3 + 3 + 6 // Adjust this based on your expectations
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}
