package main

import "fmt"

func main() {
	var N, K int

	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Ошибка чтения")
		return
	}

	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&K)
		if err != nil {
			fmt.Println("Ошибка чтения")
			return
		}

		low, high := 15, 30

		for j := 0; j < K; j++ {
			var operation string
			var T int

			_, err := fmt.Scan(&operation, &T)
			if err != nil {
				fmt.Println("Ошибка чтения")
				return
			}

			switch operation {
			case ">=":
				if T > low {
					low = T
				}
			case "<=":
				if T < high {
					high = T
				}
			default:
				fmt.Println("Неизвестная операция: ", operation)
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
