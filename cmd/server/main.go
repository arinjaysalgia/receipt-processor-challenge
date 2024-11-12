package main

import (
	"receipt-processor-challenge/pkg/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080") // Starts the server on port 8080
}
