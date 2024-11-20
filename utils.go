package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// helpers
func checkIfBetween(target PTime) bool {
	start, _ := time.Parse("15:04", "14:00")
	end, _ := time.Parse("15:04", "16:00")

	targetTime := time.Time(target)

	if targetTime.After(start) && targetTime.Before(end) {
		return true
	} else {
		return false
	}
}

func cleanForOnlyAlphaNum(str string) string {

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(str, "")
}

func checkIfZeroCents(total string) bool {
	length := len(total)
	if length < 2 {
		return false
	}

	lastTwo := total[length-2:]
	if lastTwo == "00" {
		return true
	} else {
		return false
	}
}

func checkIfTotalIsDivisibleByTwoFive(total string) bool {
	totalAsFloat, _ := strconv.ParseFloat(total, 64)

	return math.Mod(totalAsFloat, 0.25) == 0
}

func checkIfDateIsOdd(date PDate) bool {
	strDate := date.String()

	lastTwo := strDate[len(strDate)-2:]
	datePart, _ := strconv.Atoi(lastTwo)
	return datePart%2 == 1
}

func calcItemPoints(items []Item) int {

	var total = 0.0
	for _, item := range items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			priceToFloat, _ := strconv.ParseFloat(item.Price, 64)
			total += math.Ceil(priceToFloat * 0.2)

		}
	}

	return int(total)
}

func calcTotalPoints(receipt Receipt) int {
	var totalPoints = 0
	itemsLength := len(receipt.Items)

	totalPoints += len(cleanForOnlyAlphaNum(receipt.Retailer))
	if checkIfZeroCents(receipt.Total) {
		totalPoints += 50

	}
	if checkIfTotalIsDivisibleByTwoFive(receipt.Total) {
		totalPoints += 25

	}
	if itemsLength > 0 {
		totalPoints += (itemsLength / 2) * 5

		totalPoints += calcItemPoints(receipt.Items)

	}
	if checkIfDateIsOdd(receipt.PurchaseDate) {
		totalPoints += 6

	}
	if checkIfBetween(receipt.PurchaseTime) {
		totalPoints += 10

	}
	return totalPoints
}
