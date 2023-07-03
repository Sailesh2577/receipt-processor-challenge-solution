package swagger

import (
	"net/http"
	"strings"

	"encoding/json"
	"log"

	"github.com/google/uuid"
)

// Create a struct for the JSON payload
type ReceiptPayload struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ItemPayload `json:"items"`
	Total        string        `json:"total"`
}

// Create a struct for the item in the JSON payload
type ItemPayload struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var pointsStore = map[string]int{}

// processReceiptHandler handles the /receipts/process endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt JSON from the request body and parse it into a ReceiptPayload object
	var payload ReceiptPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid receipt JSON", http.StatusBadRequest)
		return
	}

	// Convert the ReceiptPayload to the Receipt struct
	receipt := Receipt{
		Retailer:     payload.Retailer,
		PurchaseDate: payload.PurchaseDate,
		PurchaseTime: payload.PurchaseTime,
		Items:        make([]Item, len(payload.Items)),
		Total:        payload.Total,
	}

	for i, itemPayload := range payload.Items {
		receipt.Items[i] = Item{
			ShortDescription: itemPayload.ShortDescription,
			Price:            itemPayload.Price,
		}
	}

	// Create an instance of the ReceiptProcessor
	processor := NewReceiptProcessor()

	// Process the receipt using the ProcessReceipt method of the ReceiptProcessor
	points := processor.ProcessReceipt(&receipt)

	// Generate an ID for the receipt
	receiptID := uuid.New().String()

	// Create a JSON response containing the ID
	response := struct {
		ID string `json:"id"`
	}{
		ID: receiptID,
	}

	// Set the response status code and content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	pointsStore[receiptID] = points

	// Write the response JSON to the response writer
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// getPointsHandler handles the /receipts/{id}/points endpoint
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from the URL path parameter
	receiptID := strings.TrimPrefix(r.URL.Path, "/receipts/")
	receiptID = strings.TrimSuffix(receiptID, "/points")

	// Retrieve the points awarded for the receipt ID (dummy implementation)
	// Implement your logic to retrieve the points based on the receiptID
	points := getPointsForReceiptID(receiptID)

	// Create a JSON response containing the points
	response := struct {
		Points int `json:"points"`
	}{
		Points: points,
	}

	// Set the response status code and content type
	w.Header().Set("Content-Type", "application/json")

	if points != -1 {
		// Write the response JSON to the response writer
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		// Return 404 if the receipt ID is not found
		http.NotFound(w, r)
	}
}

func getPointsForReceiptID(receiptID string) int {

	// Check if the receipt ID exists in the mapping
	if points, ok := pointsStore[receiptID]; ok {
		return points // Return the corresponding points
	}

	return -1 // Return -1 if receipt ID is not found
}
