package model

import "time"

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
