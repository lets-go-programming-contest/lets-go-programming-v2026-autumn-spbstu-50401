package main

import "fmt"

func main() {
	var departments, employees int
	var err error

	_, err = fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Ошибка чтения")

		return
	}

	for range departments {
		_, err = fmt.Scan(&employees)
		if err != nil {
			fmt.Println("Ошибка чтения")

			return
		}

		low, high := 15, 30

		for range employees {
			var operation string
			var temperature int

			_, err = fmt.Scan(&operation, &temperature)
			if err != nil {
				fmt.Println("Ошибка чтения")

				return
			}

			switch operation {
			case ">=":
				if temperature > low {
					low = temperature
				}
			case "<=":
				if temperature < high {
					high = temperature
				}
			default:
				fmt.Println("Неизвестная операция:", operation)

				return
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
