package main

import (
	"encoding/json"
	"net/http"
	"receipt-api/calculator"
	"receipt-api/model"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receiptMap = make(map[string]*model.Receipt)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt model.Receipt
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

	// Generate a UUID as the receipt ID
	receiptID := uuid.New().String()
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

	points := calculator.CalculatePoints(receipt)
	response := model.PointsResponse{Points: points}

	json.NewEncoder(w).Encode(response)
}
