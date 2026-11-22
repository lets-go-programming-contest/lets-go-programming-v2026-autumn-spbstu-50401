package main

import "fmt"

func main() {
	departments := 0

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		employees := 0

		_, err := fmt.Scan(&employees)
		if err != nil {
			fmt.Println(err)

			return
		}

		minTemp := 15
		maxTemp := 30

		for range employees {
			locTemp := 0
			borderSign := ' '

			_, err := fmt.Scanf("%c= %d", &borderSign, &locTemp)
			if err != nil {
				fmt.Println(err)

				return
			}

			switch borderSign {
			case '>':
				minTemp = max(minTemp, locTemp)
			case '<':
				maxTemp = min(maxTemp, locTemp)
			default:
				fmt.Println("Incorrect format")

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
