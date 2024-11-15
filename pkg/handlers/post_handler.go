package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	m "github.com/arinjaysalgia/receipt-processor-challenge/pkg/models"
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&receipt)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Validate that the Retailer field is present
	if strings.TrimSpace(receipt.Retailer) == "" {
		http.Error(w, "The 'retailer' field is required", http.StatusBadRequest)
		return
	}

	points, err := calculatePoints(receipt)
	if err != nil {
		http.Error(w, "Error processing receipt points", http.StatusInternalServerError)
		return
	}

	s := generateReceiptId()
	fmt.Println(s)
	rec := m.Receipt{
		ID:     s,
		Points: points,
	}
	m.Receipts[rec.ID] = rec

	response := map[string]string{"Id": s}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generateReceiptId() string {
	id := uuid.New()
	return id.String()
}

func calculatePoints(receipt Receipt) (int64, error) {
	points := int64(0)

	// Rule 1: One point for every alphanumeric character in the retailer name
	alphanumericRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += int64(len(alphanumericRegex.FindAllString(receipt.Retailer, -1)))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, err
	}
	if total == math.Floor(total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += int64((len(receipt.Items) / 2) * 5)

	// Rule 5: Points based on item description length
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, err
			}
			points += int64(int(math.Ceil(price * 0.2)))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is between 2:00pm and 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() == 0) {
		points += 10
	}

	return points, nil
}
