package main

import "fmt"

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid department count")
		return
	}

	for range departments {

		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Invalid employee count")
			return
		}

		minTemp := 15
		maxTemp := 30

		for range employees {
			var operator string
			var temperature int

			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				fmt.Println("Invalid input temperature data")

				return
			}

			switch operator {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			default:
				fmt.Println("Incorrect input data")

				return
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
