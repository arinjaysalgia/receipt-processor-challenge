package handlers

import (
	"encoding/json"
	"net/http"

	m "github.com/arinjaysalgia/receipt-processor-challenge/pkg/models"
	"github.com/gorilla/mux"
)

func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, exists := m.Receipts[id]
	if !exists {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	response := struct {
		Points int64 `json:"points"`
	}{
		Points: receipt.Points,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
