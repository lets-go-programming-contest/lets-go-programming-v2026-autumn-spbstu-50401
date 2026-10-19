package main

import "fmt"

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Error reading department count:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Error reading employee count:", err)

			return
		}

		lower, upper := 15, 30

		for range employeesCount {
			var (
				sign        string
				temperature int
			)

			if _, err := fmt.Scan(&sign, &temperature); err != nil {
				fmt.Println("Error reading temperature requirement:", err)

				return
			}

			switch sign {
			case ">=":
				if temperature > lower {
					lower = temperature
				}
			case "<=":
				if temperature < upper {
					upper = temperature
				}
			default:
				fmt.Printf("Incorrect input data format: unknown operator %q\n", sign)

				return
			}

			if lower <= upper {
				fmt.Println(lower)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
