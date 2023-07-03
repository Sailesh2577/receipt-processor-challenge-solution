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

	points += rp.calculateRetailerPoints(receipt.Retailer)
	points += rp.calculateRoundDollarAmountPoints(receipt.Total)
	points += rp.calculateMultipleOfQuarterPoints(receipt.Total)
	points += rp.calculateItemPoints(receipt.Items)
	points += rp.calculateTrimmedLengthPoints(receipt.Items)
	points += rp.calculateOddDayPoints(receipt.PurchaseDate)
	points += rp.calculatePurchaseTimePoints(receipt.PurchaseTime)

	return points
}

// calculateRetailerPoints calculates the points earned based on the retailer name
func (rp *ReceiptProcessor) calculateRetailerPoints(retailer string) int {
	return len(strings.ReplaceAll(retailer, " ", ""))
}

// calculateRoundDollarAmountPoints calculates the points earned if the total is a round dollar amount
func (rp *ReceiptProcessor) calculateRoundDollarAmountPoints(total string) int {
	if rp.isRoundDollarAmount(total) {
		return 50
	}
	return 0
}

// calculateMultipleOfQuarterPoints calculates the points earned if the total is a multiple of 0.25
func (rp *ReceiptProcessor) calculateMultipleOfQuarterPoints(total string) int {
	if rp.isMultipleOfQuarter(total) {
		return 25
	}
	return 0
}

// calculateItemPoints calculates the points earned based on the number of items on the receipt
func (rp *ReceiptProcessor) calculateItemPoints(items []Item) int {
	return len(items) / 2 * 5
}

// calculateTrimmedLengthPoints calculates the points earned based on the trimmed length of item descriptions
func (rp *ReceiptProcessor) calculateTrimmedLengthPoints(items []Item) int {
	points := 0
	for _, item := range items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price := rp.parsePrice(item.Price)
			points += int(price * 0.2)
		}
	}
	return points
}

// calculateOddDayPoints calculates the points earned if the day in the purchase date is odd
func (rp *ReceiptProcessor) calculateOddDayPoints(purchaseDate string) int {
	parsedDate, err := time.Parse("2006-01-02", purchaseDate)
	if err == nil && parsedDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

// calculatePurchaseTimePoints calculates the points earned based on the time of purchase
func (rp *ReceiptProcessor) calculatePurchaseTimePoints(purchaseTime string) int {
	parsedTime, err := time.Parse("15:04", purchaseTime)
	if err == nil && parsedTime.After(time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC)) && parsedTime.Before(time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)) {
		return 10
	}
	return 0
}

// isRoundDollarAmount checks if the total is a round dollar amount with no cents
func (rp *ReceiptProcessor) isRoundDollarAmount(total string) bool {
	return strings.HasSuffix(total, ".00")
}

// isMultipleOfQuarter checks if the total is a multiple of 0.25
func (rp *ReceiptProcessor) isMultipleOfQuarter(total string) bool {
	price := rp.parsePrice(total)
	return int(price*100)%25 == 0
}

// parsePrice parses the price string and returns the price as a float64
func (rp *ReceiptProcessor) parsePrice(price string) float64 {
	price = strings.ReplaceAll(price, "$", "")
	parsedPrice, _ := strconv.ParseFloat(price, 64)
	return parsedPrice
}
