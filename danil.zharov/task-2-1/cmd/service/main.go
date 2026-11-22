package main

import "fmt"

func main() {
	var departmentsCount int

	_, err := fmt.Scan(&departmentsCount)
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departmentsCount {
		var staffCount int

		_, err = fmt.Scan(&staffCount)
		if err != nil {
			fmt.Println(err)

			return
		}

		minTemp := 15
		maxTemp := 30

		for range staffCount {
			var (
				desiredTemp int
				sign        string
			)

			_, err = fmt.Scan(&sign, &desiredTemp)
			if err != nil {
				fmt.Println(err)

				return
			}

			switch sign {
			case "<=":
				maxTemp = min(maxTemp, desiredTemp)
			case ">=":
				minTemp = max(minTemp, desiredTemp)
			default:
				fmt.Println("Incorrect sign")

				return
			}

			if maxTemp >= minTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
