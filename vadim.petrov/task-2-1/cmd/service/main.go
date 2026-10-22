package main

import "fmt"

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid read count of departaments(", err, ")")

		return
	}

	for range departmentsCount {
		var peopleCount int

		if _, err := fmt.Scan(&peopleCount); err != nil {
			fmt.Println("Invalid read count of people(", err, ")")

			return
		}

		downLine, topLine := 15, 30

		for range peopleCount {
			var operator string

			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid read count of operator(", err, ")")

				return
			}

			var numTemperature int

			if _, err := fmt.Scan(&numTemperature); err != nil {
				fmt.Println("Invalid read count of temperature(", err, ")")

				return
			}

			switch operator {
			case ">=":
				downLine = numTemperature
			case "<=":
				topLine = numTemperature
			default:
				fmt.Println("Invalid operator")

				return
			}

			if topLine >= downLine {
				fmt.Println(downLine)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
