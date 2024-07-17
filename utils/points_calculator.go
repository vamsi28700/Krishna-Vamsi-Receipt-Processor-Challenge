package utils

import (
	"math"
	"strconv"
	"strings"
	"time"
	"regexp"

	"receipt-processor/models"
)

var retailerNamePattern = regexp.MustCompile("^[\\w\\s\\-&]+$")

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1 point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// If the trimmed length of the item description is a multiple of 3
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	time, _ := time.Parse("15:04", receipt.PurchaseTime)
	if time.Hour() == 14 || (time.Hour() == 15 && time.Minute() == 0) {
		points += 10
	}

	return points
}
