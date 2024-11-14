package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/arinjaysalgia/receipt-processor-challenge/pkg/handlers"
	m "github.com/arinjaysalgia/receipt-processor-challenge/pkg/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	helperfFUnc()

	r := mux.NewRouter()
	r.HandleFunc("/receipts/{id}/points", h.GetReceiptPoints).Methods("GET")
	r.HandleFunc("/receipts/process", h.ProcessReceiptHandler).Methods("POST")
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func helperfFUnc() {

	//To remove later
	s := generateReceiptId()
	fmt.Println(s)
	r0 := m.Receipt{
		ID:     s,
		Points: 329,
	}
	s = generateReceiptId()
	fmt.Println(s)
	r1 := m.Receipt{
		ID:     s,
		Points: 32,
	}
	s = generateReceiptId()
	fmt.Println(s)
	r2 := m.Receipt{
		ID:     s,
		Points: 45,
	}
	m.Receipts[r1.ID] = r1
	m.Receipts[r2.ID] = r2
	m.Receipts[r0.ID] = r0
	//Till here

}

func generateReceiptId() string {
	id := uuid.New()
	return id.String()
}
