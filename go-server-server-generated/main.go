package main

import (
	"log"
	"net/http"

	sw "github.com/Sailesh2577/receipt-processor-challenge-solution/go-server-server-generated/go"
)

func main() {
	log.Printf("Server started")

	// Initialize the router from the generated code
	router := sw.NewRouter()

	// Add your custom routes and their corresponding handlers
	// router.HandleFunc("/receipts/process", processReceiptHandler)
	// router.HandleFunc("/receipts/{id}/points", getPointsHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
