package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/arinjaysalgia/receipt-processor-challenge/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/{id}/points", h.GetReceiptPoints).Methods("GET")
	r.HandleFunc("/receipts/process", h.ProcessReceiptHandler).Methods("POST")
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
