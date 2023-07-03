package swagger

import (
	"strconv"
	"strings"
	"time"
)

// ReceiptProcessor represents the receipt processing logic
type ReceiptProcessor struct{}

// NewReceiptProcessor creates a new instance of ReceiptProcessor
func NewReceiptProcessor() *ReceiptProcessor {
	return &ReceiptProcessor{}
}

// ProcessReceipt processes the receipt and returns the number of points earned
func (rp *ReceiptProcessor) ProcessReceipt(receipt *Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if isRoundDollarAmount(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += len(receipt.Items) / 2 * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price := parsePrice(item.Price)
			points += int(price * 0.2)
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.After(time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC)) && purchaseTime.Before(time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)) {
		points += 10
	}

	return points
}

// isRoundDollarAmount checks if the total is a round dollar amount with no cents
func isRoundDollarAmount(total string) bool {
	return strings.HasSuffix(total, ".00")
}

// isMultipleOfQuarter checks if the total is a multiple of 0.25
func isMultipleOfQuarter(total string) bool {
	price := parsePrice(total)
	return int(price*100)%25 == 0
}

// parsePrice parses the price string and returns the price as a float64
func parsePrice(price string) float64 {
	price = strings.ReplaceAll(price, "$", "")
	parsedPrice, _ := strconv.ParseFloat(price, 64)
	return parsedPrice
}
