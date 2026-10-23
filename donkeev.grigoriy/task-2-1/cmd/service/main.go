package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Input reading error:", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Input reading error:", err)

			return
		}

		minTemp := 15
		maxTemp := 30

		for range employeeCount {
			var (
				operation string
				temp      int
			)

			_, err := fmt.Scan(&operation, &temp)
			if err != nil {
				fmt.Println("Input reading error:", err)

				return
			}

			switch operation {
			case ">=":
				if temp > minTemp {
					minTemp = temp
				}
			case "<=":
				if temp < maxTemp {
					maxTemp = temp
				}
			default:
				fmt.Println("Unknown operation:", operation)

				return
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
