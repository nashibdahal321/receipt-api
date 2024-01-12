package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gorilla/mux"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer      string    `json:"retailer"`
	PurchaseDate  string    `json:"purchaseDate"`
	PurchaseTime  string    `json:"purchaseTime"`
	Items         []Item    `json:"items"`
	Total         string    `json:"total"`
	ProcessedTime time.Time // To store the processing time
}

type PointsResponse struct {
	Points int `json:"points"`
}

var receiptMap = make(map[string]*Receipt)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Validate item prices
	for _, item := range receipt.Items {
		if item.Price == "" {
			http.Error(w, "Item price cannot be blank", http.StatusBadRequest)
			return
		}
	}

	receiptID := fmt.Sprintf("%x", time.Now().UnixNano())
	receiptMap[receiptID] = &receipt
	receiptMap[receiptID].ProcessedTime = time.Now()

	response := map[string]string{"id": receiptID}
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	receiptID := params["id"]

	receipt, found := receiptMap[receiptID]
	if !found {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)
	response := PointsResponse{Points: points}

	json.NewEncoder(w).Encode(response)
}

func calculatePoints(receipt *Receipt) int {
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
