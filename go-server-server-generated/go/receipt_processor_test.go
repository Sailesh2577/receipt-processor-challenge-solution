package swagger

import (
	"testing"
)

func TestReceiptProcessor_ProcessReceipt(t *testing.T) {
	processor := NewReceiptProcessor()

	receipt := &Receipt{
		Retailer:     "Example Retailer",
		PurchaseDate: "2023-07-01",
		PurchaseTime: "15:30",
		Items: []Item{
			{ShortDescription: "Item 1", Price: "$10.00"},
			{ShortDescription: "Item 2", Price: "$20.00"},
			{ShortDescription: "Item 3", Price: "$30.00"},
		},
		Total: "$60.00",
	}

	expectedPoints := 113 // Calculate the expected points based on the rules manually

	points := processor.ProcessReceipt(receipt)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_CalculateRetailerPoints(t *testing.T) {
	processor := NewReceiptProcessor()

	retailer := "Example Retailer"
	expectedPoints := 15 // Number of alphanumeric characters in the retailer name

	points := processor.calculateRetailerPoints(retailer)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_IsRoundDollarAmount(t *testing.T) {
	processor := NewReceiptProcessor()

	total := "$60.00"
	expectedResult := true // Total is a round dollar amount

	result := processor.isRoundDollarAmount(total)
	if result != expectedResult {
		t.Errorf("Expected result: %t, got: %t", expectedResult, result)
	}
}

func TestReceiptProcessor_CalculateRoundDollarAmountPoints(t *testing.T) {
	processor := NewReceiptProcessor()

	total := "$60.00"
	expectedPoints := 50 // Total is a round dollar amount

	points := processor.calculateRoundDollarAmountPoints(total)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_CalculateMultipleOfQuarterPoints(t *testing.T) {
	processor := NewReceiptProcessor()

	total := "$7.50"
	expectedPoints := 25 // Total is a multiple of 0.25

	points := processor.calculateMultipleOfQuarterPoints(total)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_CalculateItemPoints(t *testing.T) {
	processor := NewReceiptProcessor()

	items := []Item{
		{ShortDescription: "Item 1", Price: "$10.00"},
		{ShortDescription: "Item 2", Price: "$20.00"},
		{ShortDescription: "Item 3", Price: "$30.00"},
	}
	expectedPoints := 5 // Two items on the receipt

	points := processor.calculateItemPoints(items)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_CalculateTrimmedLengthPoints(t *testing.T) {
	processor := NewReceiptProcessor()

	items := []Item{
		{ShortDescription: "Item 1", Price: "$10.00"},
		{ShortDescription: "Item 2", Price: "$20.00"},
		{ShortDescription: "Item 3", Price: "$30.00"},
		{ShortDescription: "Item 4", Price: "$40.00"},
	}
	expectedPoints := 20 // Two items have trimmed length divisible by 3

	points := processor.calculateTrimmedLengthPoints(items)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_CalculateOddDayPoints(t *testing.T) {
	processor := NewReceiptProcessor()

	purchaseDate := "2023-07-01"
	expectedPoints := 6 // Day is odd

	points := processor.calculateOddDayPoints(purchaseDate)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_CalculatePurchaseTimePoints(t *testing.T) {
	processor := NewReceiptProcessor()

	purchaseTime := "15:30"
	expectedPoints := 0 // Time is between 2:00pm and 4:00pm

	points := processor.calculatePurchaseTimePoints(purchaseTime)
	if points != expectedPoints {
		t.Errorf("Expected points: %d, got: %d", expectedPoints, points)
	}
}

func TestReceiptProcessor_IsMultipleOfQuarter(t *testing.T) {
	processor := NewReceiptProcessor()

	total := "$7.50"
	expectedResult := true // Total is a multiple of 0.25

	result := processor.isMultipleOfQuarter(total)
	if result != expectedResult {
		t.Errorf("Expected result: %t, got: %t", expectedResult, result)
	}
}

func TestReceiptProcessor_ParsePrice(t *testing.T) {
	processor := NewReceiptProcessor()

	price := "$10.00"
	expectedParsedPrice := 10.0

	parsedPrice := processor.parsePrice(price)
	if parsedPrice != expectedParsedPrice {
		t.Errorf("Expected parsed price: %.2f, got: %.2f", expectedParsedPrice, parsedPrice)
	}
}
