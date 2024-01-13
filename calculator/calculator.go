package calculator

import (
	"math"
	"receipt-api/model"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(receipt *model.Receipt) int {
	// One point for every alphanumeric character in the retailer name
	trimmedRetailerName := strings.TrimSpace(receipt.Retailer)

	// Check if the trimmed name contains alphanumeric characters
	var alphanumericCount int
	for _, char := range trimmedRetailerName {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			alphanumericCount++
		}
	}
	points := alphanumericCount

	// Check for round dollar amount
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == float64(int(total)) {
		points += 50
	}

	// Check for multiple of 0.25
	if int(total*100)%25 == 0 {
		points += 25
	}

	// Check for odd purchase day
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Check for time range (2:00pm to 4:00pm)
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.After(time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)) && purchaseTime.Before(time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)) {
		points += 10
	}

	// 5 points for every two items on the receipt
	points += len(receipt.Items) / 2 * 5

	// Additional points based on item description length
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	return points
}
