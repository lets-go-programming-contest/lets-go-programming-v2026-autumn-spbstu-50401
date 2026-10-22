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

			var numTemperature int

			if _, err := fmt.Scan(&operator, &numTemperature); err != nil {
				fmt.Println("Invalid read operator or temperature(", err, ")")

				return
			}

			switch operator {
			case ">=":
				if numTemperature > downLine {
					downLine = numTemperature
				}
			case "<=":
				if numTemperature < topLine {
					topLine = numTemperature
				}
			default:
				fmt.Println("Invalid operator.")

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
