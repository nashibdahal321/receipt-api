// main/handler_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestProcessReceipt(t *testing.T) {
	// Create a sample receipt for testing
	receiptJSON := []byte(`{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
			{"shortDescription": "Knorr Creamy Chicken", "price": "1.26"},
			{"shortDescription": "Doritos Nacho Cheese", "price": "3.35"},
			{"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00"}
		],
		"total": "35.35"
	}`)

	// Create a request with the sample receipt
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(receiptJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	http.HandlerFunc(ProcessReceipt).ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ProcessReceipt returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetPoints(t *testing.T) {
	// Create a sample receipt for testing
	receiptJSON := []byte(`{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
			{"shortDescription": "Knorr Creamy Chicken", "price": "1.26"},
			{"shortDescription": "Doritos Nacho Cheese", "price": "3.35"},
			{"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00"}
		],
		"total": "35.35"
	}`)

	// Create a request with the sample receipt
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(receiptJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	http.HandlerFunc(ProcessReceipt).ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ProcessReceipt returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Extract the receipt ID from the response body
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	receiptID := response["id"]

	// Now, create a request for GetPoints using the obtained receipt ID
	req, err = http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use mux.SetURLVars to set URL variables
	vars := map[string]string{"id": receiptID}
	req = mux.SetURLVars(req, vars)

	// Create a new response recorder
	rr = httptest.NewRecorder()

	// Call the GetPoints handler function
	http.HandlerFunc(GetPoints).ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetPoints returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
