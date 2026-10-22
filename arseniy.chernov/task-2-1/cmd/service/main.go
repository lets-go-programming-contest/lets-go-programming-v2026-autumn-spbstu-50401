package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println(err)

		return
	}

	var employeeCount int
	for range departmentCount {
		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println(err)

			return
		}

		minT := 15
		maxT := 30

		var (
			temperature int
			operation   string
		)

		for range employeeCount {
			_, err := fmt.Scanf("%s %d", &operation, &temperature)
			if err != nil {
				fmt.Println(err)

				return
			}

			switch operation {
			case ">=":
				minT = max(minT, temperature)
			case "<=":
				maxT = min(maxT, temperature)
			default:
				fmt.Println("invalid operation")

				return
			}

			if minT > maxT {
				fmt.Println(-1)
			} else {
				fmt.Println(minT)
			}
		}
	}
}
