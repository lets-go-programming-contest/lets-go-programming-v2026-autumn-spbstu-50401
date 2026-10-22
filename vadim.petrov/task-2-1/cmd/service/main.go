package main

import "fmt"

func main() {
	var departatmentsCount int

	if _, err := fmt.Scanln(&departatmentsCount); err != nil {
		fmt.Println("Invalid read count of departaments(", err, ")")
		return
	}

	for range departatmentsCount {
		var peopleCount int

		if _, err := fmt.Scanln(&peopleCount); err != nil {
			fmt.Println("Invalid read count of people(", err, ")")
			return
		}

		var downLine, topLine int = 15, 30

		for range peopleCount {
			var operator string

			if _, err := fmt.Scanln(&operator); err != nil {
				fmt.Println("Invalid read count of operator(", err, ")")
				return
			}

			var numTemperature int

			if _, err := fmt.Scanln(&numTemperature); err != nil {
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

			if topLine > downLine {
				fmt.Println(downLine)
			} else {
				fmt.Println(-1)
			}
		}

	}
}
