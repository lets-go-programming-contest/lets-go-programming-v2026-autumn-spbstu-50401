package main

import (
	"fmt"
	"os"
)

type Climate struct {
	minTemp int
	maxTemp int
}

func updateTemperature(operation string, desiredTemp int, climate *Climate) bool {
	switch operation {
	case ">=":
		if desiredTemp > climate.maxTemp {
			return false
		}

		if desiredTemp > climate.minTemp {
			climate.minTemp = desiredTemp
		}
	case "<=":
		if desiredTemp < climate.minTemp {
			return false
		}

		if desiredTemp < climate.maxTemp {
			climate.maxTemp = desiredTemp
		}
	default:
		return false
	}

	return true
}

func main() {
	var countOfChapters int

	_, err := fmt.Scan(&countOfChapters)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	for range countOfChapters {
		var tempClimate Climate

		tempClimate.minTemp = 15
		tempClimate.maxTemp = 30

		var countOfEmployees int

		_, err := fmt.Scan(&countOfEmployees)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}

		flag := true

		for range countOfEmployees {
			var operation string

			_, err := fmt.Scan(&operation)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
			}

			var desiredTemp int

			_, err = fmt.Scan(&desiredTemp)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
			}

			if flag {
				flag = updateTemperature(operation, desiredTemp, &tempClimate)
			}

			if flag {
				fmt.Println(tempClimate.minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
