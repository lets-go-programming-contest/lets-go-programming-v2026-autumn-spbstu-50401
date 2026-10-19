package main

import (
	"fmt"
)

func updateTemperatureRange(operation string, number int, minTemp *int, maxTemp *int) bool {
	switch operation {
	case ">=":
		if number > *minTemp {
			*minTemp = number
		}
	case "<=":
		if number < *maxTemp {
			*maxTemp = number
		}
	default:
		fmt.Println("Invalid operation")

		return false
	}

	return true
}

func processTestCase() bool {
	var requestsCount int

	_, err := fmt.Scan(&requestsCount)
	if err != nil || requestsCount <= 0 {
		fmt.Println("Invalid input k")

		return false
	}

	minTemp := 15
	maxTemp := 30

	for range requestsCount {
		var operation string

		_, err = fmt.Scan(&operation)
		if err != nil {
			fmt.Println("Invalid input operation")

			return false
		}

		var number int

		_, err = fmt.Scan(&number)
		if err != nil || number < 15 || number > 30 {
			fmt.Println("Invalid input number")

			return false
		}

		if !updateTemperatureRange(operation, number, &minTemp, &maxTemp) {
			return false
		}

		if minTemp > maxTemp {
			fmt.Println("-1")
		} else {
			fmt.Println(minTemp)
		}
	}

	return true
}

func main() {
	var testCasesCount int

	_, err := fmt.Scan(&testCasesCount)
	if err != nil || testCasesCount <= 0 {
		fmt.Println("Invalid input n")

		return
	}

	for range testCasesCount {
		if !processTestCase() {
			return
		}
	}
}
