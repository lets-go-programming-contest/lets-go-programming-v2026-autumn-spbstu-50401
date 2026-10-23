package main

import "fmt"

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			return
		}

		lowerBound := minTemperature
		upperBound := maxTemperature

		for range employeeCount {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				return
			}

			switch operation {
			case ">=":
				if temperature > lowerBound {
					lowerBound = temperature
				}
			case "<=":
				if temperature < upperBound {
					upperBound = temperature
				}
			}

			if lowerBound <= upperBound {
				fmt.Println(lowerBound)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
